package uctx

import (
	"context"
	"github.com/zhanjunjie2019/clover/global/consts"
	"gorm.io/gorm"
)

func SetAppDB(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, consts.CtxGormDBVar, db)
}

func GetAppDB(ctx context.Context) *gorm.DB {
	value := ctx.Value(consts.CtxGormDBVar)
	if value != nil {
		db, ok := value.(*gorm.DB)
		if ok {
			return db
		}
	}
	return nil
}
