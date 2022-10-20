package defs

import (
	"context"
	"gorm.io/gorm"
)

type IRepo interface {
	AutoMigrate(ctx context.Context, db *gorm.DB) error
}
