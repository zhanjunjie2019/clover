package consumer

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/nsqio/go-nsq"
	"github.com/zhanjunjie2019/clover/global/confs"
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

type Server struct {
	Consumers     []defs.IConsumer                        `allimpls:""`
	OpenTelemetry opentelemetry.OpenTelemetryIOCInterface `singleton:""`
}

func (s *Server) ConsumersStart() error {
	nsqConfig := confs.GetGlobalConfig().NsqConfig
	if nsqConfig.Enabled == 1 {
		svcConf := confs.GetServerConfig().SvcConf
		config := nsq.NewConfig()
		for _, consumer := range s.Consumers {
			con, err := nsq.NewConsumer(consumer.GetTopic(), svcConf.SvcName, config)
			if err != nil {
				return err
			}
			con.AddHandler(newMessageHandler(consumer, s.OpenTelemetry))
			err = con.ConnectToNSQLookupd(nsqConfig.ConsumerAddr)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func newMessageHandler(consumer defs.IConsumer, provider opentelemetry.OpenTelemetryIOCInterface) nsq.Handler {
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
		zap.String("msgBody", string(pb.Body)),
	)
	err = m.consumer.HandleMessage(ctx, layout, pb.Body)
	if err != nil {
		layout.Error("消息监听错误"+err.Error(), zap.Error(err))
	}
	return err
}
