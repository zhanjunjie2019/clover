package app

import (
	"context"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/errs"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/domain/biserrs"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/domain/gateway"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/domain/model"
	_ "github.com/zhanjunjie2019/clover/starter-auth/bc/infr/gatewayimpl"
	"gorm.io/gorm"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
// +ioc:autowire:type=allimpls
// +ioc:autowire:implements=github.com/zhanjunjie2019/clover/global/defs.IAppDef

type PermissionApp struct {
	PermissionGateway gateway.IPermissionGateway `singleton:"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/gatewayimpl.PermissionGateway"`
	DB                *gorm.DB
}

func (p *PermissionApp) SetGormDB(db *gorm.DB) {
	p.DB = db
}

func (p *PermissionApp) PermissionCreate(ctx context.Context, permissionName, authCode string) (id defs.ID, err error) {
	err = p.DB.Transaction(func(tx *gorm.DB) (err error) {
		ctx = uctx.WithValueAppDB(ctx, tx)
		_, exist, err := p.PermissionGateway.FindByAuthCode(ctx, authCode)
		if err != nil {
			err = errs.ToUnifiedError(err)
			return
		}
		if exist {
			err = biserrs.PermissionAlreadyExistErrWithAuthCode(authCode)
			return
		}
		permission := model.NewPermission(0, model.PermissionValue{
			PermissionName: permissionName,
			AuthCode:       authCode,
		})
		id, err = p.PermissionGateway.Save(ctx, permission)
		if err != nil {
			err = errs.ToUnifiedError(err)
			return
		}
		return nil
	})
	return
}
