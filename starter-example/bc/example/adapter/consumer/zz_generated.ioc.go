//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package consumer

import (
	contextx "context"
	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	allimpls "github.com/alibaba/ioc-golang/extension/autowire/allimpls"
	defs "github.com/zhanjunjie2019/clover/global/defs"
)

func init() {
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &helloWorldConsumer_{}
		},
	})
	helloWorldConsumerStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &HelloWorldConsumer{}
		},
		Metadata: map[string]interface{}{
			"aop": map[string]interface{}{},
			"autowire": map[string]interface{}{
				"common": map[string]interface{}{
					"implements": []interface{}{
						new(defs.IConsumer),
					},
				},
			},
		},
	}
	allimpls.RegisterStructDescriptor(helloWorldConsumerStructDescriptor)
	var _ defs.IConsumer = &HelloWorldConsumer{}
}

type helloWorldConsumer_ struct {
	GetTopic_      func() string
	HandleMessage_ func(ctx contextx.Context, bytes []byte) error
}

func (h *helloWorldConsumer_) GetTopic() string {
	return h.GetTopic_()
}

func (h *helloWorldConsumer_) HandleMessage(ctx contextx.Context, bytes []byte) error {
	return h.HandleMessage_(ctx, bytes)
}

type HelloWorldConsumerIOCInterface interface {
	GetTopic() string
	HandleMessage(ctx contextx.Context, bytes []byte) error
}

var _helloWorldConsumerSDID string
