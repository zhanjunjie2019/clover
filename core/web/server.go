package web

import (
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	consulReg "github.com/go-micro/plugins/v4/registry/consul"
	"github.com/go-playground/validator/v10"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/zhanjunjie2019/clover/core/configs"
	"github.com/zhanjunjie2019/clover/core/sentinel"
	"github.com/zhanjunjie2019/clover/global/confs"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/errs"
	"github.com/zhanjunjie2019/clover/global/middleware"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/web"
	"net/http"
	"time"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type Server struct {
	Controllers        []defs.IController                        `allimpls:""`
	SentinelLoader     sentinel.SentinelLoaderIOCInterface       `singleton:""`
	LoggerMiddleware   middleware.LoggerMiddlewareIOCInterface   `singleton:""`
	TraceMiddleware    middleware.TraceMiddlewareIOCInterface    `singleton:""`
	SentinelMiddleware middleware.SentinelMiddlewareIOCInterface `singleton:""`
	AuthMiddleware     middleware.AuthMiddlewareIOCInterface     `singleton:""`
}

func (s *Server) RunServer() error {
	// 初始化gin
	engine := gin.New()
	// 最外层panic捕获，响应500代码
	engine.Use(gin.Recovery())
	// 服务配置
	serverConfig := confs.GetServerConfig()
	// 服务模式
	if serverConfig.SvcConf.SvcMode == confs.Release {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
		engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
	engine.GET("/health", func(c *gin.Context) {
		rsMap := gin.H{"STATUS": "UP"}
		c.JSON(http.StatusOK, rsMap)
	})
	reg := consulReg.NewRegistry(
		registry.Addrs(serverConfig.ConsulConf.ConsulAddr),
	)

	// 注册pprof
	pprof.Register(engine)

	errs.Panic(s.registRoute(engine))

	address := fmt.Sprintf(":%d", serverConfig.SvcConf.SvcPort)
	service := web.NewService(
		web.Name(serverConfig.SvcConf.SvcName),
		web.Address(address),
		web.RegisterTTL(time.Duration(serverConfig.ConsulConf.RegisterTTL)*time.Second),
		web.RegisterInterval(time.Duration(serverConfig.ConsulConf.RegisterInterval)*time.Second),
		web.Version(serverConfig.SvcConf.SvcVersion),
		web.Registry(reg),
		web.Handler(engine),
	)
	// 初始化服务
	errs.Panic(service.Init())
	return service.Run()
}

var validate = validator.New()

func (s Server) registRoute(engine *gin.Engine) error {
	sentineEnabled := confs.GetSentinelConfig().Enabled
	// 加载全局限流配置
	if sentineEnabled == 1 {
		s.SentinelLoader.AppendServerRules()
	}
	for _, c := range s.Controllers {
		// 读取接口配置
		option := c.GetOption()
		err := validate.Struct(option)
		if err != nil {
			return err
		}
		// 动态中间件
		var handlerFuncs []gin.HandlerFunc
		if sentineEnabled == 1 {
			handlerFuncs = append(handlerFuncs, s.SentinelMiddleware.MiddlewareHandlerFunc(nil))
		}
		handlerFuncs = append(handlerFuncs, s.TraceMiddleware.MiddlewareHandlerFunc(nil))
		handlerFuncs = append(handlerFuncs, s.LoggerMiddleware.MiddlewareHandlerFunc(&option))
		// 如果属于限权接口
		if len(option.AuthCodes) > 0 {
			// 资源级限流中间件
			handlerFuncs = append(handlerFuncs, s.AuthMiddleware.MiddlewareHandlerFunc(&option))
		}
		if sentineEnabled == 1 && len(option.SentinelStrategy) > 0 {
			// 接口限流规则导入
			s.SentinelLoader.AppendApiRules(option)
			// 资源级限流中间件
			handlerFuncs = append(handlerFuncs, s.SentinelMiddleware.MiddlewareHandlerFunc(&option))
		}
		// 加载自定义的业务中间件
		middlewares := option.Middlewares
		if middlewares != nil && len(middlewares) > 0 {
			for _, mid := range middlewares {
				handlerFuncs = append(handlerFuncs, mid.MiddlewareHandlerFunc(&option))
			}
		}
		handlerFuncs = append(handlerFuncs, c.Handle)
		// 注册接口路由
		engine.Handle(option.HttpMethod, option.RelativePath, handlerFuncs...)
	}
	if sentineEnabled == 1 {
		return s.SentinelLoader.LoadSentinelRules()
	}
	return nil
}
