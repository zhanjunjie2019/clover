package app

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zhanjunjie2019/clover/global/consts"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/errs"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"github.com/zhanjunjie2019/clover/global/utils"
	"github.com/zhanjunjie2019/clover/share/auth/dto"
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

func (t *TenantApp) TenantCreate(ctx context.Context, tenantID, tenantName, redirect string, accessTTL, refreshTTL uint64) (tid, secretKey string, err error) {
	ctx = uctx.WithValueAppDB(ctx, t.DB)
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
		TenantID:              tid,
		TenantName:            tenantName,
		RedirectUrl:           redirect,
		AccessTokenTimeLimit:  accessTTL,
		RefreshTokenTimeLimit: refreshTTL,
	})
	tenant.GenerateSecretKey()
	_, err = t.TenantGateway.Save(ctx, tenant)
	if err != nil {
		err = errs.ToUnifiedError(err)
		return
	}
	// 发布广播通知租户初始化
	err = t.TenantGateway.PublishInitEvent(ctx, dto.TenantInitEventDTO{
		TenantID:   tid,
		TenantName: tenantName,
	})
	if err != nil {
		err = errs.ToUnifiedError(err)
		return
	}
	return tenant.FullValue().TenantID, tenant.FullValue().SecretKey, nil
}

func (t *TenantApp) TenantInit(ctx context.Context, tenantID string) (err error) {
	ctx = uctx.WithValueTenantAndAppDB(ctx, tenantID, t.DB)
	return t.TenantGateway.TenantTablesManualMigrate(ctx)
}

func (t *TenantApp) TenantTokenCreate(ctx context.Context, tenantID, secretKey string, accessTokenExpTime, refreshTokenExpTime int64) (accessToken, refreshToken string, accessTokenExpirationTime, refreshTokenExpirationTime int64, err error) {
	ctx = uctx.WithValueAppDB(ctx, t.DB)
	tenant, exist, err := t.TenantGateway.FindByTenantID(ctx, tenantID)
	if err != nil {
		err = errs.ToUnifiedError(err)
		return
	}
	if !exist {
		err = biserrs.TenantDoesNotExistErr
		return
	}
	if !tenant.VerifySecretKey(secretKey) {
		err = biserrs.TenantValidationFailedErr
		return
	}
	// 构建accessToken
	accessExp := jwt.NumericDate{}
	if accessTokenExpTime > 0 {
		accessExp.Time = time.Unix(accessTokenExpTime, 0)
		accessTokenExpirationTime = accessTokenExpTime
	} else {
		accessExp.Time = time.Now().Add(time.Second * time.Duration(tenant.FullValue().AccessTokenTimeLimit))
		accessTokenExpirationTime = accessExp.Time.Unix()
	}
	accessJwtClaims := defs.JwtClaims{
		TenantID: tenantID,
		Auths:    []string{consts.AdminAuth},
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        utils.UUID(),
			ExpiresAt: &accessExp,
		},
	}
	accessToken, err = uctx.CreateJwtClaimsToken(accessJwtClaims)
	if err != nil {
		err = errs.ToUnifiedError(err)
		return
	}
	// 构建refreshToken
	refreshExp := jwt.NumericDate{}
	if refreshTokenExpTime > 0 {
		refreshExp.Time = time.Unix(refreshTokenExpTime, 0)
		refreshTokenExpirationTime = refreshTokenExpTime
	} else {
		refreshExp.Time = time.Now().Add(time.Second * time.Duration(tenant.FullValue().RefreshTokenTimeLimit))
		refreshTokenExpirationTime = refreshExp.Time.Unix()
	}
	refreshJwtClaims := defs.JwtClaims{
		TenantID: tenantID,
		Auths:    []string{consts.RefreshAdminAuth},
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        utils.UUID(),
			ExpiresAt: &refreshExp,
		},
	}
	refreshToken, err = uctx.CreateJwtClaimsToken(refreshJwtClaims)
	if err != nil {
		err = errs.ToUnifiedError(err)
		return
	}
	return
}
