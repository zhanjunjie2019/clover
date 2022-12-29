package grpc

import (
	"fmt"
	"github.com/go-micro/plugins/v4/registry/consul"
	grpcsvr "github.com/go-micro/plugins/v4/server/grpc"
	"github.com/zhanjunjie2019/clover/core/sentinel"
	"github.com/zhanjunjie2019/clover/global/confs"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/errs"
	"github.com/zhanjunjie2019/clover/global/middleware"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/server"
	"time"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type Server struct {
	Handlers           []defs.IGrpcServiceHandler                `allimpls:""`
	SentinelLoader     sentinel.SentinelLoaderIOCInterface       `singleton:""`
	LoggerMiddleware   middleware.LoggerMiddlewareIOCInterface   `singleton:""`
	TraceMiddleware    middleware.TraceMiddlewareIOCInterface    `singleton:""`
	SentinelMiddleware middleware.SentinelMiddlewareIOCInterface `singleton:""`
	AuthMiddleware     middleware.AuthMiddlewareIOCInterface     `singleton:""`
}

func (s *Server) RunServer() error {
	// 服务配置
	serverConfig := confs.GetServerConfig()
	// grpc服务
	grpcServer := grpcsvr.NewServer(
		server.Name(serverConfig.SvcConf.SvcName),
		server.Version(serverConfig.SvcConf.SvcVersion+"-grpc"),
		server.Address(fmt.Sprintf(":%d", serverConfig.SvcConf.Grpc.Port)),
		server.Registry(consul.NewRegistry(
			registry.Addrs(serverConfig.ConsulConf.ConsulAddr),
		)),
		server.RegisterTTL(time.Duration(serverConfig.ConsulConf.RegisterTTL)*time.Second),
		server.RegisterInterval(time.Duration(serverConfig.ConsulConf.RegisterInterval)*time.Second),
	)
	// grpc服务
	grpcService := micro.NewService(
		micro.Server(grpcServer),
	)
	// 初始化服务，与通用中间件
	grpcService.Init(
		micro.WrapHandler(s.SentinelMiddleware.MiddlewareWrapHandler()),
		micro.WrapHandler(s.TraceMiddleware.MiddlewareWrapHandler()),
		micro.WrapHandler(s.LoggerMiddleware.MiddlewareWrapHandler()),
	)

	// 注册核心业务Grpc接口
	errs.Panic(s.registHandler(grpcService.Server()))
	return grpcService.Run()
}

func (s Server) registHandler(sv server.Server) error {
	sentineEnabled := confs.GetSentinelConfig().Enabled
	// 加载全局限流配置
	if sentineEnabled == 1 {
		s.SentinelLoader.AppendServerRules()
	}
	// 注册gRPC核心业务接口
	for _, handler := range s.Handlers {
		// 注册事件
		errs.Panic(handler.GrpcRegister(sv))
	}
	if sentineEnabled == 1 {
		return s.SentinelLoader.LoadSentinelRules()
	}
	return nil
}
