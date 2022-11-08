package defs

import (
	"context"
)

type IRepo interface {
	AutoMigrate(ctx context.Context) error
}
