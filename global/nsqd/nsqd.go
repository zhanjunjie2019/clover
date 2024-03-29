package nsqd

import (
	"context"
	"github.com/gogo/protobuf/proto"
	"github.com/nsqio/go-nsq"
	"github.com/zhanjunjie2019/clover/global/confs"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/nsqd/protobuf"
	"github.com/zhanjunjie2019/clover/global/opentelemetry"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"github.com/zhanjunjie2019/clover/global/uorm"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"runtime"
	"sync"
	"time"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type NsqProducer struct {
	addr          string
	rw            sync.RWMutex
	producer      *nsq.Producer
	OpenTelemetry opentelemetry.OpenTelemetryIOCInterface `singleton:""`
}

func (n *NsqProducer) CreatePublisher(producerAddr string) (bool, error) {
	n.rw.Lock()
	defer n.rw.Unlock()
	if n.addr != producerAddr {
		producer, err := nsq.NewProducer(producerAddr, nsq.NewConfig())
		if err != nil {
			return false, err
		}
		if n.producer != nil {
			n.producer.Stop()
		}
		n.addr = producerAddr
		n.producer = producer
		return true, nil
	}
	return false, nil
}

func (n *NsqProducer) Publish(ctx context.Context, topic string, body []byte) (err error) {
	n.rw.RLock()
	defer n.rw.RUnlock()
	_, span := n.OpenTelemetry.Start(ctx, "Producer "+topic)
	defer func() {
		if err == nil {
			span.End()
		} else {
			span.RecordError(err)
		}
	}()
	msg := protobuf.NsqMessage{
		TenantID: uctx.GetTenantID(ctx),
		Body:     body,
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

func NewMessageHandler(consumer defs.IConsumer, provider opentelemetry.OpenTelemetryIOCInterface, dbFactory uorm.DBFactoryIOCInterface) nsq.Handler {
	return &messageHandler{
		provider:  provider,
		consumer:  consumer,
		dbFactory: dbFactory,
	}
}

type messageHandler struct {
	provider  opentelemetry.OpenTelemetryIOCInterface
	consumer  defs.IConsumer
	dbFactory uorm.DBFactoryIOCInterface
}

func (m *messageHandler) HandleMessage(message *nsq.Message) (err error) {
	svcConf := confs.GetServerConfig().SvcConf
	layout := defs.NewLogLayout(zapcore.InfoLevel, svcConf.SvcMode.Uint8(), svcConf.SvcName, svcConf.SvcNum, svcConf.SvcVersion)

	var pb protobuf.NsqMessage
	err = proto.Unmarshal(message.Body, &pb)
	if err != nil {
		layout.Error("消息监听错误"+err.Error(), zap.Error(err))
		return err
	}
	// 链路上下文中传递日志输出器
	ctx := uctx.WithValueLogLayout(context.Background(), layout)
	// 数据库实例对象
	db := m.dbFactory.GetDB()
	if db != nil {
		ctx = uctx.WithValueAppDB(ctx, db)
	}
	// 链路上下文中传递租户ID
	if len(pb.TenantID) > 0 {
		ctx = uctx.WithValueTenantID(ctx, pb.TenantID)
		layout.AppendLogsFields(zap.String("tenantID", pb.TenantID))
	}
	ctx = m.provider.GetCtx(ctx, pb.TraceId, pb.TraceSpanID)
	ctx, span := m.provider.Start(ctx, "Consumer "+m.consumer.GetTopic())
	start := time.Now()
	defer func() {
		// 异常日志处理
		if recoverErr := recover(); recoverErr != nil {
			buf := make([]byte, 1<<12)
			runtime.Stack(buf, false)
			layout.Error("消息监听故障：" + string(buf))
			span.RecordError(err)
		} else if err != nil {
			layout.Error("消息监听故障："+err.Error(), zap.Error(err))
			span.RecordError(err)
		} else {
			span.End()
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

	err = m.consumer.HandleMessage(ctx, pb.Body)
	if err != nil {
		layout.Error("消息监听错误"+err.Error(), zap.Error(err))
	}
	return err
}
