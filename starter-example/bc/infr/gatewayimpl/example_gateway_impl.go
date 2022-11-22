package gatewayimpl

import (
	"context"
	"github.com/gogo/protobuf/proto"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/nsqd"
	"github.com/zhanjunjie2019/clover/share/example/protobuf"
	"github.com/zhanjunjie2019/clover/share/example/topic"
	"github.com/zhanjunjie2019/clover/starter-example/bc/domain/model"
	"github.com/zhanjunjie2019/clover/starter-example/bc/infr/gatewayimpl/convs"
	"github.com/zhanjunjie2019/clover/starter-example/bc/infr/repo"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type ExampleGateway struct {
	Entity1Repo repo.ExampleEntity1RepoIOCInterface `singleton:""`
	Entity2Repo repo.ExampleEntity2RepoIOCInterface `singleton:""`
	NsqProducer nsqd.NsqProducerIOCInterface        `singleton:""`
}

func (e *ExampleGateway) SaveExample1(ctx context.Context, entity1 model.Entity1) (id defs.ID, err error) {
	entity1PO, entity2PO := convs.Entity1DOToPOWithEntity2(entity1)
	id, err = e.Entity1Repo.Save(ctx, entity1PO)
	if err != nil {
		return
	}
	entity2PO.Entity1ID = id
	_, err = e.Entity2Repo.Save(ctx, entity2PO)
	return
}

func (e *ExampleGateway) PublishEventMessage(ctx context.Context, dto protobuf.ExampleDTO) error {
	bs, err := proto.Marshal(&dto)
	if err != nil {
		return err
	}
	return e.NsqProducer.Publish(ctx, topic.ExampleTopic, bs)
}
