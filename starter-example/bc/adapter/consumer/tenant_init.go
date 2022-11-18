package consumer

import (
	"context"
	"encoding/json"
	"github.com/zhanjunjie2019/clover/share/auth/dto"
	"github.com/zhanjunjie2019/clover/share/auth/topic"
)

// +ioc:autowire=true
// +ioc:autowire:type=allimpls
// +ioc:autowire:implements=github.com/zhanjunjie2019/clover/global/defs.IConsumer

type TenantInitConsumer struct {
}

func (t *TenantInitConsumer) GetTopic() string {
	return topic.TenantInitTopic
}

func (t *TenantInitConsumer) HandleMessage(ctx context.Context, bytes []byte) error {
	var dto dto.TenantInitEventDTO
	err := json.Unmarshal(bytes, &dto)
	if err != nil {
		return err
	}
	return nil
}
