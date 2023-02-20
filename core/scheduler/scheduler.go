package scheduler

import (
	"context"
	"github.com/zhanjunjie2019/clover/global/confs"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/opentelemetry"
	"github.com/zhanjunjie2019/clover/global/redisc"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"github.com/zhanjunjie2019/clover/global/uorm"
	"github.com/zhanjunjie2019/clover/global/utils"
	"go.opentelemetry.io/otel/trace"
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
	DBFactory     uorm.DBFactoryIOCInterface              `singleton:""`
	timer         utils.Timer
}

func (s *Server) RegistryServer() error {
	s.timer = utils.NewTimer()
	for _, sc := range s.Schedulers {
		err := s.startScheduler(sc)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Server) startScheduler(scheduler defs.IScheduler) error {
	svcConf := confs.GetServerConfig().SvcConf
	_, err := s.timer.AddTaskByFunc(
		scheduler.GetTaskTypeCode(),
		scheduler.GetSpec(),
		func() {
			if scheduler.GetLockDuration() > 0 {
				// 竞争分布式并发锁
				ok, _ := redisc.RedisConcurrentLockInTime(context.Background(), s.RedisClient, svcConf.SvcName+"."+scheduler.GetTaskTypeCode(), scheduler.GetLockDuration())
				if !ok {
					return
				}
			}
			// 链路上下文中传递日志输出器
			layout := defs.NewLogLayout(zapcore.InfoLevel, svcConf.SvcMode.Uint8(), svcConf.SvcName, svcConf.SvcNum, svcConf.SvcVersion)
			ctx := uctx.WithValueLogLayout(context.Background(), layout)

			// 数据库实例对象
			db := s.DBFactory.GetDB()
			if db != nil {
				ctx = uctx.WithValueAppDB(ctx, db)
			}

			start := time.Now()

			if scheduler.LoggerEnable() {
				// 开启一个根级span
				var span trace.Span
				ctx, span = s.OpenTelemetry.Start(ctx, "Scheduler "+scheduler.GetTaskTypeCode())
				defer span.End()
			}

			err := scheduler.RunTask(ctx)
			if err != nil {
				layout.Error(err.Error(), zap.Error(err))
				layout.AppendLogsFields(
					zap.Time("startTime", start),
					zap.Time("endTime", time.Now()),
					zap.String("taskTypeCode", scheduler.GetTaskTypeCode()),
				)
				layout.Println()
			} else if scheduler.LoggerEnable() {
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
