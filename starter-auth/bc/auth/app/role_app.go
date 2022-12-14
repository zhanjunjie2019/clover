package app

import (
	"context"
	"github.com/samber/lo"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/errs"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/app/cmd"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/domain/biserrs"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/domain/gateway"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/domain/model"
	_ "github.com/zhanjunjie2019/clover/starter-auth/bc/auth/infr/gatewayimpl"
	"gorm.io/gorm"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
// +ioc:autowire:type=allimpls
// +ioc:autowire:implements=github.com/zhanjunjie2019/clover/global/defs.IAppDef

type RoleApp struct {
	RoleGateway       gateway.IRoleGateway       `singleton:"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/infr/gatewayimpl.RoleGateway"`
	PermissionGateway gateway.IPermissionGateway `singleton:"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/infr/gatewayimpl.PermissionGateway"`
	DB                *gorm.DB
}

func (r *RoleApp) SetGormDB(db *gorm.DB) {
	r.DB = db
}

func (r *RoleApp) RoleCreate(ctx context.Context, c cmd.RoleCreateCmd) (rs cmd.RoleCreateResult, err error) {
	err = r.DB.Transaction(func(tx *gorm.DB) (err error) {
		ctx = uctx.WithValueAppDB(ctx, tx)
		roleCodes := lo.Map(c.Roles, func(item cmd.RoleInfo, index int) string {
			return item.RoleCode
		})
		roles, err := r.RoleGateway.ListByRoleCodes(ctx, roleCodes)
		if err != nil {
			err = errs.ToUnifiedError(err)
			return
		}
		if len(roles) > 0 {
			roleCodes = lo.Map(roles, func(item model.Role, index int) string {
				return item.FullValue().RoleCode
			})
			err = biserrs.RoleAlreadyExistsErrWithRoleCode(roleCodes...)
			return
		}
		for _, roleInfo := range c.Roles {
			role := model.NewRole(0, model.RoleValue{
				RoleName: roleInfo.RoleName,
				RoleCode: roleInfo.RoleCode,
			})
			var id defs.ID
			id, err = r.RoleGateway.SaveSingle(ctx, role)
			if err != nil {
				err = errs.ToUnifiedError(err)
				return
			}
			rs.PermissionIDs = append(rs.PermissionIDs, id)
		}
		return nil
	})
	return
}

func (r *RoleApp) RolePermissionAssignment(ctx context.Context, c cmd.RolePermissionAssignmentCmd) (rs cmd.RolePermissionAssignmentResult, err error) {
	err = r.DB.Transaction(func(tx *gorm.DB) (err error) {
		ctx = uctx.WithValueAppDB(ctx, tx)
		// ?????????????????????
		var (
			role  model.Role
			exist bool
		)
		if c.RoleID == 0 {
			role, exist, err = r.RoleGateway.FindByRoleCode(ctx, c.RoleCode)
		} else {
			role, exist, err = r.RoleGateway.FindByID(ctx, c.RoleID)
		}
		if err != nil {
			err = errs.ToUnifiedError(err)
			return
		}
		if !exist {
			err = biserrs.RoleDoesNotExistErr
			return
		}
		// ????????????????????????
		permissions, err := r.PermissionGateway.ListByAuthCodes(ctx, c.AuthCodes)
		if err != nil {
			err = errs.ToUnifiedError(err)
			return
		}
		var pvs []model.PermissionValue
		// ?????????????????????????????????????????????????????????
		for _, authCode := range c.AuthCodes {
			var pv *model.PermissionValue
			for pi := range permissions {
				val := permissions[pi].FullValue()
				if val.AuthCode == authCode {
					pv = &val
					break
				}
			}
			if pv == nil {
				err = biserrs.PermissionDoesNotExistErrWithAuthCode(authCode)
				return
			}
			pvs = append(pvs, *pv)
		}
		// ?????????????????????
		role.SetPermissionValues(pvs)
		// ???????????????
		rs.RoleID, err = r.RoleGateway.SaveWithPermission(ctx, role)
		if err != nil {
			err = errs.ToUnifiedError(err)
			return
		}
		return nil
	})
	return
}
