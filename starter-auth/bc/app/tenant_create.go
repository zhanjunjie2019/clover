package app

import (
	"context"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/errs"
	"github.com/zhanjunjie2019/clover/global/utils"
	"github.com/zhanjunjie2019/clover/share/auth/dto"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/domain/biserrs"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/domain/gateway"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/domain/model"
	_ "github.com/zhanjunjie2019/clover/starter-auth/bc/infr/gatewayimpl"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type TenantCreateApp struct {
	TenantGateway gateway.ITenantGateway `singleton:"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/gatewayimpl.TenantGateway"`
}

func (t *TenantCreateApp) TenantCreate(ctx context.Context, layout *defs.LogLayout, tenantID, tenantName string) (tid, secretKey string, err error) {
	if len(tenantID) == 0 {
		tid = utils.UUID()
	} else {
		tid = tenantID
	}
	_, exist, err := t.TenantGateway.FindByTenantID(ctx, tid)
	if err != nil {
		err = errs.ToUnifiedError(err)
		return
	}
	if exist {
		err = biserrs.TenantAlreadyExistErr
		return
	}
	tenant := model.NewTenant(0, model.TenantValue{
		TenantID:   tenantID,
		TenantName: tenantName,
	})
	tenant.GenerateSecretKey()
	err = t.TenantGateway.Save(ctx, tenant)
	if err != nil {
		err = errs.ToUnifiedError(err)
		return
	}
	// 发布广播通知租户初始化
	err = t.TenantGateway.PublishInitEvent(ctx, dto.TenantInitEventDTO{
		TenantID:   tenantID,
		TenantName: tenantName,
	})
	if err != nil {
		err = errs.ToUnifiedError(err)
		return
	}
	return tenant.FullValue().TenantID, tenant.FullValue().SecretKey, nil
}
