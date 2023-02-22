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
	_ "github.com/zhanjunjie2019/clover/starter-auth/bc"
	_ "github.com/zhanjunjie2019/clover/starter-auth/docs"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @Title       clover-auth-api
// @Version     v1.0.0
// @Description clover-auth接口文档
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
	// 加载本地配置，必须
	errs.Panic(s.ConfigDefines.LoadAllConfigByLocal())
	// 初始化日志组件，必须
	errs.Panic(logs.InitLogger())
	// 加载远程配置，非必须
	errs.Panic(s.ConfigDefines.LoadAllConfigByConsul())
	// 初始化遥感链路追踪，非必须
	errs.Panic(s.OpenTelemetry.InitProvider())
	// 初始化数据库连接，非必须
	errs.Panic(s.RepoDBFactory.Initialization())
	// 启动定时任务调度，非必须
	errs.Panic(s.SchedulerServer.RegistryServer())
	// 注册NSQ消息队列监听，非必须
	errs.Panic(s.ConsumerServer.RegistryServer())
	// 启动HTTP服务，与GRPC至少启一个
	return s.WebServer.RunServer()
}
