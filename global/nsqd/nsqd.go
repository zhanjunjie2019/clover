package nsqd

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/nsqio/go-nsq"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/nsqd/protobuf"
	"github.com/zhanjunjie2019/clover/global/opentelemetry"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"runtime"
	"time"
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

func (n *NsqProducer) Publish(ctx context.Context, topic string, body []byte) (err error) {
	_, span := n.OpenTelemetry.Start(ctx, "Producer "+topic)
	defer func() {
		if err == nil {
			span.End()
		} else {
			span.RecordError(err)
		}
	}()
	msg := protobuf.NsqMessage{
		Body: body,
	}
	sc := span.SpanContext()
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

func NewMessageHandler(consumer defs.IConsumer, provider opentelemetry.OpenTelemetryIOCInterface) nsq.Handler {
	return &messageHandler{
		provider: provider,
		consumer: consumer,
	}
}

type messageHandler struct {
	provider opentelemetry.OpenTelemetryIOCInterface
	consumer defs.IConsumer
}

func (m *messageHandler) HandleMessage(message *nsq.Message) error {
	layout := defs.NewLogLayout(zapcore.InfoLevel)
	var pb protobuf.NsqMessage
	err := proto.Unmarshal(message.Body, &pb)
	if err != nil {
		layout.Error("消息监听错误"+err.Error(), zap.Error(err))
		return err
	}
	ctx := m.provider.GetCtx(context.Background(), pb.TraceId, pb.TraceSpanID)
	ctx, span := m.provider.Start(ctx, "Consumer "+m.consumer.GetTopic())
	defer span.End()
	start := time.Now()
	defer func() {
		// 异常日志处理
		if recoverErr := recover(); recoverErr != nil {
			buf := make([]byte, 1<<12)
			runtime.Stack(buf, false)
			layout.Error("消息监听故障：" + string(buf))
		}
		layout.AppendLogsFields(
			zap.Int64("rt", time.Since(start).Milliseconds()),
		)
		layout.Println()
	}()
	layout.AppendLogsFields(
		zap.String("topic", m.consumer.GetTopic()),
		zap.String("traceID", pb.TraceId),
		zap.String("msgBody", string(pb.Body)),
	)
	err = m.consumer.HandleMessage(ctx, layout, pb.Body)
	if err != nil {
		layout.Error("消息监听错误"+err.Error(), zap.Error(err))
	}
	return err
}
