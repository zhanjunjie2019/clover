package gateway

import (
	"context"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/share/example/protobuf"
	"github.com/zhanjunjie2019/clover/starter-example/bc/domain/model"
)

type IExampleGateway interface {
	SaveExample1(ctx context.Context, entity1 model.Entity1) (defs.ID, error)
	PublishEventMessage(ctx context.Context, dto protobuf.ExampleDTO) error
}
