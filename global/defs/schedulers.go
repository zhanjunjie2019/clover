package defs

import (
	contextx "context"
	"time"
)

type IScheduler interface {
	// GetTaskTypeCode 定时任务编码，用于防止并发验证，要求全局唯一
	GetTaskTypeCode() string
	// GetSpec corn表达式
	GetSpec() string
	// GetLockDuration 并发限制间隔，多久时间内限制重复并发
	GetLockDuration() time.Duration
	// RunTask 指定核心任务
	RunTask(contextx.Context, *LogLayout) error
}
