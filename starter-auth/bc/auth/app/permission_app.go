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
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type PermissionApp struct {
	PermissionGateway gateway.IPermissionGateway `singleton:"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/infr/gatewayimpl.PermissionGateway"`
}

func (p *PermissionApp) PermissionCreate(ctx context.Context, c cmd.PermissionCreateCmd) (rs cmd.PermissionCreateResult, err error) {
	err = uctx.AppTransaction(ctx, func(ctx context.Context) (err error) {
		authCodes := lo.Map(c.Permissions, func(item cmd.PermissionInfo, index int) string {
			return item.AuthCode
		})
		permissions, err := p.PermissionGateway.ListByAuthCodes(ctx, authCodes)
		if err != nil {
			err = errs.ToUnifiedError(err)
			return
		}
		if len(permissions) > 0 {
			authCodes = lo.Map(permissions, func(item model.Permission, index int) string {
				return item.FullValue().AuthCode
			})
			err = biserrs.PermissionAlreadyExistErrWithAuthCode(authCodes...)
			return
		}
		for _, permissionInfo := range c.Permissions {
			permission := model.NewPermission(0, model.PermissionValue{
				PermissionName: permissionInfo.PermissionName,
				AuthCode:       permissionInfo.AuthCode,
			})
			var id defs.ID
			id, err = p.PermissionGateway.Save(ctx, permission)
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
