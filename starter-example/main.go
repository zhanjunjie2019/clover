package main

import (
	"github.com/alibaba/ioc-golang"
	"github.com/zhanjunjie2019/clover/core/consumer"
	"github.com/zhanjunjie2019/clover/core/repo"
	"github.com/zhanjunjie2019/clover/core/scheduler"
	"github.com/zhanjunjie2019/clover/core/web"
	"github.com/zhanjunjie2019/clover/global/config"
	"github.com/zhanjunjie2019/clover/global/errs"
	"github.com/zhanjunjie2019/clover/global/logs"
	"github.com/zhanjunjie2019/clover/global/opentelemetry"
	_ "github.com/zhanjunjie2019/clover/starter-example/bc/adapter/consumer"
	_ "github.com/zhanjunjie2019/clover/starter-example/bc/adapter/scheduler"
	_ "github.com/zhanjunjie2019/clover/starter-example/docs"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @Title       clover-example-api
// @Version     v1.0.0
// @Description clover-example接口文档
// @BasePath    /example
// @SecurityDefinitions.Apikey ApiKeyAuth
// @In header
// @Name C-Token

func main() {
	// 加载依赖注入
	errs.Panic(ioc.Load())

	// 获取starter实例
	starter, err := GetStarterSingleton()
	errs.Panic(err)

	// 运行starter
	errs.Panic(starter.Run())
}

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type Starter struct {
	SchedulerServer scheduler.ServerIOCInterface            `singleton:""`
	ConsumerServer  consumer.ServerIOCInterface             `singleton:""`
	WebServer       web.ServerIOCInterface                  `singleton:""`
	RepoDBFactory   repo.RepoDBFactoryIOCInterface          `singleton:""`
	OpenTelemetry   opentelemetry.OpenTelemetryIOCInterface `singleton:""`
	ConfigDefines   config.ConfigDefinesIOCInterface        `singleton:""`
}

func (s *Starter) Run() error {
	// 加载本地配置
	errs.Panic(s.ConfigDefines.LoadAllConfigByLocal())
	// 初始化日志组件
	errs.Panic(logs.InitLogger())
	// 加载远程配置
	errs.Panic(s.ConfigDefines.LoadAllConfigByConsul())
	// 初始化遥感链路追踪
	errs.Panic(s.OpenTelemetry.InitProvider())
	// 初始化数据库连接
	errs.Panic(s.RepoDBFactory.Initialization())
	// 启动定时任务调度
	errs.Panic(s.SchedulerServer.SchedulersStart())
	// 启动NSQ消息队列监听
	errs.Panic(s.ConsumerServer.ConsumersStart())
	// 启动HTTP请求监听
	return s.WebServer.RunServer()
}
