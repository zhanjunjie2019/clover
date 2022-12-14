package app

import (
	"context"
	"github.com/zhanjunjie2019/clover/global/consts"
	"github.com/zhanjunjie2019/clover/global/errs"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"github.com/zhanjunjie2019/clover/global/utils"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/domain/biserrs"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/domain/gateway"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/domain/model"
	_ "github.com/zhanjunjie2019/clover/starter-auth/bc/infr/gatewayimpl"
	"gorm.io/gorm"
	"time"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
// +ioc:autowire:type=allimpls
// +ioc:autowire:implements=github.com/zhanjunjie2019/clover/global/defs.IAppDef

type TenantApp struct {
	TenantGateway gateway.ITenantGateway `singleton:"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/gatewayimpl.TenantGateway"`
	DB            *gorm.DB
}

func (t *TenantApp) SetGormDB(db *gorm.DB) {
	t.DB = db
}

func (t *TenantApp) TenantCreate(ctx context.Context, tenantID, tenantName, redirect string, accessTTL uint64) (tid, secretKey string, err error) {
	err = t.DB.Transaction(func(tx *gorm.DB) (err error) {
		ctx = uctx.WithValueAppDB(ctx, tx)
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
			err = biserrs.TenantAlreadyExistErrWithTenantID(tenantID)
			return
		}
		tenant := model.NewTenant(0, model.TenantValue{
			TenantID:             tid,
			TenantName:           tenantName,
			RedirectUrl:          redirect,
			AccessTokenTimeLimit: accessTTL,
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
		return nil
	})
	return
}

func (t *TenantApp) TenantInit(ctx context.Context, tenantID string) (err error) {
	ctx = uctx.WithValueTenantAndAppDB(ctx, tenantID, t.DB)
	return t.TenantGateway.TenantTablesManualMigrate(ctx)
}

func (t *TenantApp) TenantTokenCreate(ctx context.Context, tenantID, secretKey string, accessTokenExpTime int64) (accessToken string, accessTokenExpirationTime int64, err error) {
	ctx = uctx.WithValueAppDB(ctx, t.DB)
	tenant, exist, err := t.TenantGateway.FindByTenantID(ctx, tenantID)
	if err != nil {
		err = errs.ToUnifiedError(err)
		return
	}
	if !exist {
		err = biserrs.TenantDoesNotExistErrWithTenantID(tenantID)
		return
	}
	if !tenant.VerifySecretKey(secretKey) {
		err = biserrs.TenantValidationFailedErr
		return
	}
	if accessTokenExpTime == 0 {
		accessTokenExpirationTime = time.Now().Add(time.Second * time.Duration(tenant.FullValue().AccessTokenTimeLimit)).Unix()
	} else {
		accessTokenExpirationTime = accessTokenExpTime
	}
	accessToken, err = uctx.NewJwtClaimsToken(tenantID, 0, "", []string{consts.AdminAuth}, accessTokenExpirationTime)
	if err != nil {
		err = errs.ToUnifiedError(err)
		return
	}
	return
}
