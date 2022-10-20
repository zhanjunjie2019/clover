package repo

import (
	"context"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/repo/po"
	"gorm.io/gorm"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
// +ioc:autowire:type=allimpls
// +ioc:autowire:implements=github.com/zhanjunjie2019/clover/global/defs.IRepo

type TenantRepo struct {
	db *gorm.DB
}

func (t *TenantRepo) AutoMigrate(ctx context.Context, db *gorm.DB) error {
	t.db = db
	return t.db.WithContext(ctx).AutoMigrate(po.Tenant{})
}

func (t *TenantRepo) FindByTenantID(ctx context.Context, tenantID string) (tenantPO po.Tenant, exist bool, err error) {
	err = t.db.WithContext(ctx).Where("tenant_id=?", tenantID).Last(&tenantPO).Error
	if err == nil {
		exist = true
	} else if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return
}

func (t *TenantRepo) Save(ctx context.Context, tenantPO po.Tenant) error {
	return t.db.WithContext(ctx).Save(&tenantPO).Error
}
