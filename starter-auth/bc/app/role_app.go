package app

import (
	"context"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/errs"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/domain/biserrs"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/domain/gateway"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/domain/model"
	"gorm.io/gorm"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
// +ioc:autowire:type=allimpls
// +ioc:autowire:implements=github.com/zhanjunjie2019/clover/global/defs.IAppDef

type RoleApp struct {
	RoleGateway gateway.IRoleGateway `singleton:"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/gatewayimpl.RoleGateway"`
	DB          *gorm.DB
}

func (r *RoleApp) SetGormDB(db *gorm.DB) {
	r.DB = db
}

func (r *RoleApp) RoleCreate(ctx context.Context, roleName, roleCode string) (id defs.ID, err error) {
	err = r.DB.Transaction(func(tx *gorm.DB) (err error) {
		ctx = uctx.WithValueAppDB(ctx, tx)
		_, exist, err := r.RoleGateway.FindByRoleCode(ctx, roleCode)
		if err != nil {
			err = errs.ToUnifiedError(err)
			return
		}
		if exist {
			err = biserrs.RoleAlreadyExistsErr
			return
		}
		role := model.NewRole(0, model.RoleValue{
			RoleName: roleName,
			RoleCode: roleCode,
		})
		id, err = r.RoleGateway.Save(ctx, role)
		if err != nil {
			err = errs.ToUnifiedError(err)
			return
		}
		return nil
	})
	return
}
