package main

import (
	"github.com/alibaba/ioc-golang"
	"github.com/zhanjunjie2019/clover/core/web"
	"github.com/zhanjunjie2019/clover/global/config"
	"github.com/zhanjunjie2019/clover/global/errs"
	"github.com/zhanjunjie2019/clover/global/logs"
	_ "github.com/zhanjunjie2019/clover/starter-generation/docs"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @Title       clover-generation-api
// @Version     v1.0.0
// @Description clover-generation接口文档
// @SecurityDefinitions.Apikey ApiKeyAuth
// @In header
// @Name C-Token

func main() {
	// 加载依赖注入，必须
	errs.Panic(ioc.Load())

	// 获取starter实例，必须
	starter, err := GetStarterSingleton()
	errs.Panic(err)

	// 运行starter，必须
	errs.Panic(starter.Run())
}

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type Starter struct {
	WebServer     web.ServerIOCInterface           `singleton:""`
	ConfigDefines config.ConfigDefinesIOCInterface `singleton:""`
}

func (s *Starter) Run() error {
	// 加载本地配置，必须
	errs.Panic(s.ConfigDefines.LoadAllConfigByLocal())
	// 初始化日志组件，必须
	errs.Panic(logs.InitLogger())
	// 启动HTTP请求监听，必须
	return s.WebServer.RunServer()
}
