package app

import (
	"context"
	"github.com/zhanjunjie2019/clover/global/defs"
	"gorm.io/gorm"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
// +ioc:autowire:type=allimpls
// +ioc:autowire:implements=github.com/zhanjunjie2019/clover/global/defs.IAppDef

type PermissionApp struct {
	DB *gorm.DB
}

func (p *PermissionApp) SetGormDB(db *gorm.DB) {
	p.DB = db
}

func (p *PermissionApp) PermissionCreate(ctx context.Context, permissionName, authCode string) (id defs.ID, err error) {
	err = p.DB.Transaction(func(tx *gorm.DB) (err error) {

		return nil
	})
	return
}
