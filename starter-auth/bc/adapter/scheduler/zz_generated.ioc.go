//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package scheduler

import (
	contextx "context"
	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	allimpls "github.com/alibaba/ioc-golang/extension/autowire/allimpls"
	"github.com/zhanjunjie2019/clover/global/defs"
	timex "time"
)

func init() {
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &secondScheduler_{}
		},
	})
	secondSchedulerStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &SecondScheduler{}
		},
		Metadata: map[string]interface{}{
			"aop": map[string]interface{}{},
			"autowire": map[string]interface{}{
				"common": map[string]interface{}{
					"implements": []interface{}{
						new(defs.IScheduler),
					},
				},
			},
		},
	}
	allimpls.RegisterStructDescriptor(secondSchedulerStructDescriptor)
}

type secondScheduler_ struct {
	GetTaskTypeCode_ func() string
	GetSpec_         func() string
	GetLockDuration_ func() timex.Duration
	RunTask_         func(ctx contextx.Context, layout *defs.LogLayout) error
}

func (s *secondScheduler_) GetTaskTypeCode() string {
	return s.GetTaskTypeCode_()
}

func (s *secondScheduler_) GetSpec() string {
	return s.GetSpec_()
}

func (s *secondScheduler_) GetLockDuration() timex.Duration {
	return s.GetLockDuration_()
}

func (s *secondScheduler_) RunTask(ctx contextx.Context, layout *defs.LogLayout) error {
	return s.RunTask_(ctx, layout)
}

type SecondSchedulerIOCInterface interface {
	GetTaskTypeCode() string
	GetSpec() string
	GetLockDuration() timex.Duration
	RunTask(ctx contextx.Context, layout *defs.LogLayout) error
}

var _secondSchedulerSDID string
