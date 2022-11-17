package uctx

import (
	"context"
	"github.com/zhanjunjie2019/clover/global/consts"
	"gorm.io/gorm"
)

func WithValueAppDB(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, consts.CtxGormDBVar, db)
}

func WithValueTenantAndAppDB(ctx context.Context, tenantID string, db *gorm.DB) context.Context {
	return context.WithValue(context.WithValue(ctx, consts.CtxTenantIDVar, tenantID), consts.CtxGormDBVar, db)
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

func GetAppDBWithCtx(ctx context.Context) *gorm.DB {
	return GetAppDB(ctx).WithContext(ctx)
}
