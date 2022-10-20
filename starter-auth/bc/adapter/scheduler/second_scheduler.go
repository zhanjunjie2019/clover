package scheduler

import (
	contextx "context"
	"fmt"
	"github.com/zhanjunjie2019/clover/global/defs"
	"time"
)

// +ioc:autowire=true
// +ioc:autowire:type=allimpls
// +ioc:autowire:implements=github.com/zhanjunjie2019/clover/global/defs.IScheduler

type SecondScheduler struct {
}

func (s *SecondScheduler) GetTaskTypeCode() string {
	return "auth.SecondScheduler"
}

func (s *SecondScheduler) GetSpec() string {
	return "* * * * * ?"
}

func (s *SecondScheduler) GetLockDuration() time.Duration {
	return 500 * time.Microsecond
}

func (s *SecondScheduler) RunTask(ctx contextx.Context, layout *defs.LogLayout) error {
	fmt.Println("auth.SecondScheduler")
	return nil
}
