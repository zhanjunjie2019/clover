package app

import (
	"context"
	"github.com/zhanjunjie2019/clover/starter-generation/bc/generation/app/cmd"
	"github.com/zhanjunjie2019/clover/starter-generation/bc/generation/domain/gateway"
	"github.com/zhanjunjie2019/clover/starter-generation/bc/generation/domain/model"
	_ "github.com/zhanjunjie2019/clover/starter-generation/bc/generation/infr/gatewayimpl"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type GenerationApp struct {
	GenerationGateway gateway.IGenerationGateway `singleton:"github.com/zhanjunjie2019/clover/starter-generation/bc/generation/infr/gatewayimpl.GenerationGateway"`
}

func (g *GenerationApp) ModuleGen(ctx context.Context, c cmd.ModuleGenCmd) (rs cmd.ModuleGenResult, err error) {
	module := model.NewModule(0, model.ModuleValue{
		RootPackagePath:       c.RootPackagePath,
		ModuleName:            c.ModuleName,
		ServerPort:            c.ServerPort,
		EnabledConfigByConsul: c.EnabledConfigByConsul,
		EnabledOpenTelemetry:  c.EnabledOpenTelemetry,
	})
	err = g.GenerationGateway.ModuleGen(ctx, module)
	return
}
