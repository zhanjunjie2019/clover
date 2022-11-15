//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package nsqd

import (
	contextx "context"
	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	singleton "github.com/alibaba/ioc-golang/autowire/singleton"
	util "github.com/alibaba/ioc-golang/autowire/util"
)

func init() {
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &nsqProducer_{}
		},
	})
	nsqProducerStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &NsqProducer{}
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
	}
	singleton.RegisterStructDescriptor(nsqProducerStructDescriptor)
}

type nsqProducer_ struct {
	CreatePublisher_ func(producerAddr string) error
	Publish_         func(ctx contextx.Context, topic string, body []byte) error
}

func (n *nsqProducer_) CreatePublisher(producerAddr string) error {
	return n.CreatePublisher_(producerAddr)
}

func (n *nsqProducer_) Publish(ctx contextx.Context, topic string, body []byte) error {
	return n.Publish_(ctx, topic, body)
}

type NsqProducerIOCInterface interface {
	CreatePublisher(producerAddr string) error
	Publish(ctx contextx.Context, topic string, body []byte) error
}

var _nsqProducerSDID string

func GetNsqProducerSingleton() (*NsqProducer, error) {
	if _nsqProducerSDID == "" {
		_nsqProducerSDID = util.GetSDIDByStructPtr(new(NsqProducer))
	}
	i, err := singleton.GetImpl(_nsqProducerSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*NsqProducer)
	return impl, nil
}

func GetNsqProducerIOCInterfaceSingleton() (NsqProducerIOCInterface, error) {
	if _nsqProducerSDID == "" {
		_nsqProducerSDID = util.GetSDIDByStructPtr(new(NsqProducer))
	}
	i, err := singleton.GetImplWithProxy(_nsqProducerSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(NsqProducerIOCInterface)
	return impl, nil
}
