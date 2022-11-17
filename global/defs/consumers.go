package defs

import "context"

type IConsumer interface {
	GetTopic() string
	HandleMessage(context.Context, []byte) error
}
