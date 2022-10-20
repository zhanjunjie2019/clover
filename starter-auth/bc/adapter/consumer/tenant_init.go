package consumer

import (
	"context"
	"fmt"
	"github.com/zhanjunjie2019/clover/global/defs"
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

func (t *TenantInitConsumer) HandleMessage(ctx context.Context, layout *defs.LogLayout, bytes []byte) error {
	fmt.Println(string(bytes))
	return nil
}
