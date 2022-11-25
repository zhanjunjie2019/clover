package gateway

import (
	"context"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/starter-example/bc/domain/model"
)

type IExampleGateway interface {
	SaveExampleEntity(ctx context.Context, entity model.ExampleEntity) (defs.ID, error)
	PublishEventMessage(ctx context.Context, entity model.ExampleEntity) error
}
