package uctx

import (
	"context"
	"github.com/zhanjunjie2019/clover/global/consts"
	"gorm.io/gorm"
)

func WithValueAppDB(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, consts.CtxGormDBVar, db)
}

func WithValueTenant(ctx context.Context, tenantID string) context.Context {
	return context.WithValue(ctx, consts.CtxTenantIDVar, tenantID)
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

func AppTransaction(ctx context.Context, fn func(context.Context) error) error {
	return GetAppDB(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(WithValueAppDB(ctx, tx))
	})
}

func GetAppDBWithCtx(ctx context.Context) *gorm.DB {
	return GetAppDB(ctx).WithContext(ctx)
}

func GetTenantTableDBWithCtx(ctx context.Context, tablePrefile string) *gorm.DB {
	return GetAppDB(ctx).WithContext(ctx).Table(tablePrefile + "_" + GetTenantID(ctx))
}
