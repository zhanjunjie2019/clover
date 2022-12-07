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

type RoleApp struct {
	RoleGateway       gateway.IRoleGateway       `singleton:"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/gatewayimpl.RoleGateway"`
	PermissionGateway gateway.IPermissionGateway `singleton:"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/gatewayimpl.PermissionGateway"`
	DB                *gorm.DB
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
			err = biserrs.RoleAlreadyExistsErr(roleCode)
			return
		}
		role := model.NewRole(0, model.RoleValue{
			RoleName: roleName,
			RoleCode: roleCode,
		})
		id, err = r.RoleGateway.SaveSingle(ctx, role)
		if err != nil {
			err = errs.ToUnifiedError(err)
			return
		}
		return nil
	})
	return
}

func (r *RoleApp) RolePermissionAssignment(ctx context.Context, roleCode string, authCodes []string) (id defs.ID, err error) {
	err = r.DB.Transaction(func(tx *gorm.DB) (err error) {
		ctx = uctx.WithValueAppDB(ctx, tx)
		// 验证角色有效性
		role, exist, err := r.RoleGateway.FindByRoleCode(ctx, roleCode)
		if err != nil {
			err = errs.ToUnifiedError(err)
			return
		}
		if !exist {
			err = biserrs.RoleDoesNotExistErr(roleCode)
			return
		}
		// 获取资源许可列表
		permissions, err := r.PermissionGateway.ListByAuthCodes(ctx, authCodes)
		if err != nil {
			err = errs.ToUnifiedError(err)
			return
		}
		var pers []model.PermissionValue
		// 对比数据库资源与参数的，进行有效性校验
		for i := range authCodes {
			authCode := authCodes[i]
			var pv *model.PermissionValue
			for pi := range permissions {
				val := permissions[pi].FullValue()
				if val.AuthCode == authCode {
					pv = &val
					break
				}
			}
			if pv == nil {
				err = biserrs.PermissionDoesNotExistErr(authCode)
				return
			}
			pers = append(pers, *pv)
		}
		// 校验完成后赋值
		role.SetPermissionValues(pers)
		// 保存聚合根
		id, err = r.RoleGateway.SaveWithPermission(ctx, role)
		if err != nil {
			err = errs.ToUnifiedError(err)
			return
		}
		return nil
	})
	return
}
