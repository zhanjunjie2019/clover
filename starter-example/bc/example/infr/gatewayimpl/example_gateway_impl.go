package gatewayimpl

import (
	"context"
	"github.com/gogo/protobuf/proto"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/nsqd"
	"github.com/zhanjunjie2019/clover/share/example/protobuf"
	"github.com/zhanjunjie2019/clover/share/example/topic"
	"github.com/zhanjunjie2019/clover/starter-example/bc/example/domain/model"
	"github.com/zhanjunjie2019/clover/starter-example/bc/example/infr/gatewayimpl/convs"
	"github.com/zhanjunjie2019/clover/starter-example/bc/example/infr/repo"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type ExampleGateway struct {
	ExampleEntityRepo      repo.ExampleEntityRepoIOCInterface      `singleton:""`
	ExampleValueObjectRepo repo.ExampleValueObjectRepoIOCInterface `singleton:""`
	NsqProducer            nsqd.NsqProducerIOCInterface            `singleton:""`
}

func (e *ExampleGateway) SaveExampleEntity(ctx context.Context, entity model.ExampleEntity) (id defs.ID, err error) {
	exampleEntity, exampleValueObject := convs.EntityDOToPOWithValueObject(entity)
	id, err = e.ExampleEntityRepo.Save(ctx, exampleEntity)
	if err != nil {
		return
	}
	exampleValueObject.EntityID = id
	_, err = e.ExampleValueObjectRepo.Save(ctx, exampleValueObject)
	return
}

func (e *ExampleGateway) PublishEventMessage(ctx context.Context, entity model.ExampleEntity) error {
	dto := protobuf.ExampleDTO{
		FirstName: entity.FullValue().FirstName,
		LastName:  entity.FullValue().LastName,
	}
	bs, err := proto.Marshal(&dto)
	if err != nil {
		return err
	}
	return e.NsqProducer.Publish(ctx, topic.ExampleTopic, bs)
}
