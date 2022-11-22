package consumer

import (
	"context"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/zhanjunjie2019/clover/share/example/protobuf"
	"github.com/zhanjunjie2019/clover/share/example/topic"
)

type HelloWordConsumer struct {
}

func (h *HelloWordConsumer) GetTopic() string {
	return topic.ExampleTopic
}

func (h *HelloWordConsumer) HandleMessage(ctx context.Context, bytes []byte) error {
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
