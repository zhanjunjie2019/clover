package consumer

import (
	"github.com/nsqio/go-nsq"
	"github.com/zhanjunjie2019/clover/global/confs"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/nsqd"
	"github.com/zhanjunjie2019/clover/global/opentelemetry"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type Server struct {
	Consumers     []defs.IConsumer                        `allimpls:""`
	OpenTelemetry opentelemetry.OpenTelemetryIOCInterface `singleton:""`
}

func (s *Server) RegistryServer() error {
	nsqConfig := confs.GetGlobalConfig().NsqConfig
	if nsqConfig.Enabled == 1 {
		svcConf := confs.GetServerConfig().SvcConf
		config := nsq.NewConfig()
		for _, consumer := range s.Consumers {
			con, err := nsq.NewConsumer(consumer.GetTopic(), svcConf.SvcName, config)
			if err != nil {
				return err
			}
			con.AddHandler(nsqd.NewMessageHandler(consumer, s.OpenTelemetry))
			err = con.ConnectToNSQLookupd(nsqConfig.ConsumerAddr)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
