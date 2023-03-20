package app

import (
	"context"
	"github.com/samber/lo"
	"github.com/zhanjunjie2019/clover/global/consts"
	"github.com/zhanjunjie2019/clover/global/errs"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"github.com/zhanjunjie2019/clover/global/utils"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/app/cmd"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/domain/biserrs"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/domain/gateway"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/domain/model"
	_ "github.com/zhanjunjie2019/clover/starter-auth/bc/auth/infr/gatewayimpl"
	"time"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type TenantApp struct {
	TenantGateway gateway.ITenantGateway `singleton:"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/infr/gatewayimpl.TenantGateway"`
}

func (t *TenantApp) TenantCreate(ctx context.Context, c cmd.TenantCreateCmd) (rs cmd.TenantCreateResult, err error) {
	err = uctx.AppTransaction(ctx, func(ctx context.Context) (err error) {
		// 先筛选出带租户ID的数据，在获取租户ID查询数据库是否重复
		tenantIds := lo.Map(lo.Filter(c.Tenants, func(item cmd.TenantInfo, index int) bool {
			return len(item.TenantID) > 0
		}), func(item cmd.TenantInfo, index int) string {
			return item.TenantID
		})
		if len(tenantIds) > 0 {
			var tenants []model.Tenant
			tenants, err = t.TenantGateway.ListByTenantIDs(ctx, tenantIds)
			if err != nil {
				err = errs.ToUnifiedError(err)
				return
			}
			if len(tenants) > 0 {
				tenantIds = lo.Map(tenants, func(item model.Tenant, index int) string {
					return item.FullValue().TenantID
				})
				err = biserrs.TenantAlreadyExistErrWithTenantID(tenantIds...)
				return
			}
		}
		for _, tenantInfo := range c.Tenants {
			var (
				tid       string
				secretKey string
			)
			if len(tenantInfo.TenantID) == 0 {
				tid = utils.TinyUUID()
			} else {
				tid = tenantInfo.TenantID
			}
			tenant := model.NewTenant(0, model.TenantValue{
				TenantID:             tid,
				TenantName:           tenantInfo.TenantName,
				RedirectUrl:          tenantInfo.RedirectUrl,
				AccessTokenTimeLimit: tenantInfo.AccessTokenTimeLimit,
			})
			secretKey = tenant.GenerateSecretKey()
			_, err = t.TenantGateway.Save(ctx, tenant)
			if err != nil {
				err = errs.ToUnifiedError(err)
				return
			}
			// 发布广播通知租户初始化
			err = t.TenantGateway.PublishInitEvent(ctx, tenant)
			if err != nil {
				err = errs.ToUnifiedError(err)
				return
			}
			rs.SecretKeys = append(rs.SecretKeys, cmd.TenantSecretKey{
				TenantID:  tid,
				SecretKey: secretKey,
			})
		}
		return nil
	})
	return
}

func (t *TenantApp) TenantInit(ctx context.Context, c cmd.TenantInitCmd) (err error) {
	ctx = uctx.WithValueTenant(ctx, c.TenantID)
	return t.TenantGateway.TenantTablesManualMigrate(ctx)
}

func (t *TenantApp) TenantTokenCreate(ctx context.Context, c cmd.TenantTokenCreateCmd) (rs cmd.TenantTokenCreateResult, err error) {
	tenant, exist, err := t.TenantGateway.FindByTenantID(ctx, c.TenantID)
	if err != nil {
		err = errs.ToUnifiedError(err)
		return
	}
	if !exist {
		err = biserrs.TenantDoesNotExistErrWithTenantID(c.TenantID)
		return
	}
	if !tenant.VerifySecretKey(c.SecretKey) {
		err = biserrs.TenantValidationFailedErr
		return
	}
	if c.AccessTokenExpirationTime == 0 {
		rs.AccessTokenExpirationTime = time.Now().Unix() + int64(tenant.FullValue().AccessTokenTimeLimit)
	} else {
		rs.AccessTokenExpirationTime = c.AccessTokenExpirationTime
	}
	rs.AccessToken, err = uctx.NewJwtClaimsToken(c.TenantID, 0, "", []string{consts.AdminAuth}, rs.AccessTokenExpirationTime)
	if err != nil {
		err = errs.ToUnifiedError(err)
		return
	}
	return
}
