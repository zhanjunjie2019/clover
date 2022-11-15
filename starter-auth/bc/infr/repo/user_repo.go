package repo

import (
	"context"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/repo/po"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type UserRepo struct {
}

func (u *UserRepo) ManualMigrate(ctx context.Context) error {
	p := &po.User{TenantID: uctx.GetTenantID(ctx)}
	return uctx.GetAppDBWithCtx(ctx).Table(p.TableName()).AutoMigrate(p)
}
