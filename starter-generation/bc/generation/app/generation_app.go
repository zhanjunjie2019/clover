package app

import (
	"context"
	"github.com/zhanjunjie2019/clover/starter-generation/bc/generation/app/cmd"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type GenerationApp struct {
}

func (g *GenerationApp) ModuleGen(ctx context.Context, c cmd.ModuleGenCmd) (rs cmd.ModuleGenResult, err error) {
	return
}
