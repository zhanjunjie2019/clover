package nsqd

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/nsqio/go-nsq"
	"github.com/zhanjunjie2019/clover/global/nsqd/protobuf"
	"github.com/zhanjunjie2019/clover/global/opentelemetry"
	"go.opentelemetry.io/otel/trace"
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
	c, span := n.OpenTelemetry.Start(ctx, "Producer "+topic)
	defer span.End()
	msg := protobuf.NsqMessage{
		Body: body,
	}
	sc := trace.SpanContextFromContext(c)
	if sc.IsValid() {
		msg.TraceId = sc.TraceID().String()
		msg.TraceSpanID = sc.SpanID().String()
	}
	bytes, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	return n.producer.Publish(topic, bytes)
}
