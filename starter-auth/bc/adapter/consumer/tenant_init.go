package consumer

import (
	"context"
	"github.com/gogo/protobuf/proto"
	"github.com/zhanjunjie2019/clover/share/auth/protobuf"
	"github.com/zhanjunjie2019/clover/share/auth/topic"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/app"
)

// +ioc:autowire=true
// +ioc:autowire:type=allimpls
// +ioc:autowire:implements=github.com/zhanjunjie2019/clover/global/defs.IConsumer

type TenantInitConsumer struct {
	TenantApp app.TenantAppIOCInterface `singleton:""`
}

func (t *TenantInitConsumer) GetTopic() string {
	return topic.TenantInitTopic
}

func (t *TenantInitConsumer) HandleMessage(ctx context.Context, bytes []byte) error {
	var dto protobuf.TenantInitEventDTO
	err := proto.Unmarshal(bytes, &dto)
	if err != nil {
		return err
	}
	return t.TenantApp.TenantInit(ctx, dto.TenantID)
}
