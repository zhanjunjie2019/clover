package consumer

import (
	"context"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/zhanjunjie2019/clover/share/example/protobuf"
	"github.com/zhanjunjie2019/clover/share/example/topic"
)

// +ioc:autowire=true
// +ioc:autowire:type=allimpls
// +ioc:autowire:implements=github.com/zhanjunjie2019/clover/global/defs.IConsumer

type HelloWorldConsumer struct {
}

func (h *HelloWorldConsumer) GetTopic() string {
	return topic.ExampleTopic
}

func (h *HelloWorldConsumer) HandleMessage(ctx context.Context, bytes []byte) error {
	// json格式内容日志可见，但是包大性能低
	// proto格式内容日志不可见，但是包小性能高
	// 根据自身业务特征，选择
	var dto protobuf.ExampleDTO
	err := proto.Unmarshal(bytes, &dto)
	if err != nil {
		return err
	}
	fmt.Println(dto.FirstName + " " + dto.LastName)
	return nil
}
