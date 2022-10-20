package consumer

import (
	"context"
	"github.com/nsqio/go-nsq"
	"github.com/zhanjunjie2019/clover/global/confs"
	"github.com/zhanjunjie2019/clover/global/defs"
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
	// 开启一个根级span
	ctx, span := m.provider.Start(context.Background(), "Consumer "+m.consumer.GetTopic())
	defer span.End()
	start := time.Now()
	layout := defs.NewLogLayout(zapcore.InfoLevel)
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
		zap.String("msgBody", string(message.Body)),
	)
	err := m.consumer.HandleMessage(ctx, layout, message.Body)
	if err != nil {
		layout.Error("消息监听错误"+err.Error(), zap.Error(err))
	}
	return err
}
