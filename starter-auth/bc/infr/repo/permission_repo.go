package repo

import (
	"context"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/repo/po"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
// +ioc:autowire:type=allimpls
// +ioc:autowire:implements=github.com/zhanjunjie2019/clover/global/defs.IRepo

type PermissionRepo struct {
}

func (p *PermissionRepo) AutoMigrate(ctx context.Context) error {
	return uctx.GetAppDBWithCtx(ctx).AutoMigrate(&po.Permission{})
}
