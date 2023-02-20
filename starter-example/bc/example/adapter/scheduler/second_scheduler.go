package scheduler

import (
	contextx "context"
	"time"
)

// +ioc:autowire=true
// +ioc:autowire:type=allimpls
// +ioc:autowire:implements=github.com/zhanjunjie2019/clover/global/defs.IScheduler

type SecondScheduler struct {
}

func (s *SecondScheduler) GetTaskTypeCode() string {
	return "example.SecondScheduler"
}

func (s *SecondScheduler) GetSpec() string {
	return "0 * * * * ?"
}

func (s *SecondScheduler) GetLockDuration() time.Duration {
	return time.Second
}

func (s *SecondScheduler) RunTask(ctx contextx.Context) error {
	return nil
}

func (s *SecondScheduler) LoggerEnable() bool {
	return true
}
