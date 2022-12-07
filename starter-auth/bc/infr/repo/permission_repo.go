package repo

import (
	"context"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/repo/po"
	"gorm.io/gorm"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
// +ioc:autowire:type=allimpls
// +ioc:autowire:implements=github.com/zhanjunjie2019/clover/global/defs.IRepo

type PermissionRepo struct {
}

func (p *PermissionRepo) AutoMigrate(ctx context.Context) error {
	return uctx.GetAppDBWithCtx(ctx).AutoMigrate(po.Permission{})
}

func (p *PermissionRepo) FindByAuthCode(ctx context.Context, authCode string) (permissionPO po.Permission, exist bool, err error) {
	err = uctx.GetAppDBWithCtx(ctx).Where("auth_code=?", authCode).First(&permissionPO).Error
	if err == nil {
		exist = true
	} else if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return
}

func (p *PermissionRepo) Save(ctx context.Context, permissionPO po.Permission) (defs.ID, error) {
	err := uctx.GetAppDBWithCtx(ctx).Save(&permissionPO).Error
	return permissionPO.ID, err
}

func (p *PermissionRepo) ListByAuthCodes(ctx context.Context, authCodes []string) (permissionPOs []po.Permission, err error) {
	err = uctx.GetAppDBWithCtx(ctx).Where("auth_code IN ?", authCodes).Find(&permissionPOs).Error
	return
}
