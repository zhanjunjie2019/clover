package app

import (
	"context"
	"github.com/zhanjunjie2019/clover/global/consts"
	"github.com/zhanjunjie2019/clover/global/errs"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/domain/biserrs"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/domain/model"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/configs"
	"time"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type SadminApp struct {
}

func (s *SadminApp) SadminTokenCreate(ctx context.Context, secretKey string) (accessToken string, accessTokenExpirationTime int64, err error) {
	auperAdmin := configs.GetAuthConfig().SuperAdmin
	tenant := model.NewTenant(0, model.TenantValue{
		SecretKey:            auperAdmin.SecretKey,
		AccessTokenTimeLimit: auperAdmin.AccessTokenTimeLimit,
	})
	if !tenant.VerifySecretKey(secretKey) {
		err = biserrs.TenantValidationFailedErr
		return
	}
	accessTokenExpirationTime = time.Now().Add(time.Second * time.Duration(tenant.FullValue().AccessTokenTimeLimit)).Unix()
	accessToken, err = uctx.NewJwtClaimsToken("", 0, "", []string{consts.SAdminAuth}, accessTokenExpirationTime)
	if err != nil {
		err = errs.ToUnifiedError(err)
		return
	}
	return
}
