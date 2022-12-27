package gatewayimpl

import (
	"context"
	"github.com/zhanjunjie2019/clover/starter-generation/bc/generation/domain/model"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type GenerationGateway struct {
}

func (g *GenerationGateway) ModuleGen(ctx context.Context, module model.Module) error {
	//TODO implement me
	panic("implement me")
}
