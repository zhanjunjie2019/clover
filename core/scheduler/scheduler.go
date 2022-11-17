package scheduler

import (
	"context"
	"github.com/zhanjunjie2019/clover/global/confs"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/opentelemetry"
	"github.com/zhanjunjie2019/clover/global/redisc"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"github.com/zhanjunjie2019/clover/global/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type Server struct {
	Schedulers    []defs.IScheduler                       `allimpls:""`
	RedisClient   redisc.RedisClientIOCInterface          `singleton:""`
	OpenTelemetry opentelemetry.OpenTelemetryIOCInterface `singleton:""`
	timer         utils.Timer
}

func (s *Server) SchedulersStart() error {
	svcConf := confs.GetServerConfig().SvcConf
	s.timer = utils.NewTimer()
	for _, sc := range s.Schedulers {
		err := s.startScheduler(sc, svcConf.SvcName)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Server) startScheduler(scheduler defs.IScheduler, svcName string) error {
	_, err := s.timer.AddTaskByFunc(
		scheduler.GetTaskTypeCode(),
		scheduler.GetSpec(),
		func() {
			// 竞争分布式并发锁
			ok, _ := redisc.RedisConcurrentLockInTime(context.Background(), s.RedisClient, svcName+"."+scheduler.GetTaskTypeCode(), scheduler.GetLockDuration())
			if ok {
				// 链路上下文中传递日志输出器
				layout := defs.NewLogLayout(zapcore.InfoLevel)
				ctx := uctx.WithValueLogLayout(context.Background(), layout)

				start := time.Now()
				// 开启一个根级span
				ctx, span := s.OpenTelemetry.Start(ctx, "Scheduler "+scheduler.GetTaskTypeCode())
				defer span.End()

				err := scheduler.RunTask(ctx)
				if err != nil {
					layout.Error(err.Error(), zap.Error(err))
				}
				layout.AppendLogsFields(
					zap.Time("startTime", start),
					zap.Time("endTime", time.Now()),
					zap.String("taskTypeCode", scheduler.GetTaskTypeCode()),
				)
				layout.Println()
			}
		},
	)
	return err
}
