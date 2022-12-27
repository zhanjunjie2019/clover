package gateway

import (
	"context"
	"github.com/zhanjunjie2019/clover/starter-generation/bc/generation/domain/model"
)

type IGenerationGateway interface {
	ModuleGen(ctx context.Context, module model.Module) error
}
