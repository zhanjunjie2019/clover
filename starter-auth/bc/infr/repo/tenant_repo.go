package repo

import (
	"context"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/repo/po"
	"gorm.io/gorm"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
// +ioc:autowire:type=allimpls
// +ioc:autowire:implements=github.com/zhanjunjie2019/clover/global/defs.IRepo

type TenantRepo struct {
}

func (t *TenantRepo) AutoMigrate(ctx context.Context) error {
	return uctx.GetAppDBWithCtx(ctx).AutoMigrate(po.Tenant{})
}

func (t *TenantRepo) FindByTenantID(ctx context.Context, tenantID string) (tenantPO po.Tenant, exist bool, err error) {
	err = uctx.GetAppDBWithCtx(ctx).Where("tenant_id=?", tenantID).First(&tenantPO).Error
	if err == nil {
		exist = true
	} else if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return
}

func (t *TenantRepo) Save(ctx context.Context, tenantPO po.Tenant) error {
	return uctx.GetAppDBWithCtx(ctx).Save(&tenantPO).Error
}
