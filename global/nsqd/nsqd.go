package nsqd

import (
	"context"
	"github.com/nsqio/go-nsq"
	"github.com/zhanjunjie2019/clover/global/opentelemetry"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type NsqProducer struct {
	producer      *nsq.Producer
	OpenTelemetry opentelemetry.OpenTelemetryIOCInterface `singleton:""`
}

func (n *NsqProducer) CreatePublisher(producerAddr string) error {
	producer, err := nsq.NewProducer(producerAddr, nsq.NewConfig())
	if err != nil {
		return err
	}
	if n.producer != nil {
		n.producer.Stop()
	}
	n.producer = producer
	return nil
}

func (n *NsqProducer) Publish(ctx context.Context, topic string, body []byte) error {
	_, span := n.OpenTelemetry.Start(ctx, "Producer "+topic)
	defer span.End()
	return n.producer.Publish(topic, body)
}
