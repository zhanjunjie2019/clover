package app

import (
	"context"
	"github.com/zhanjunjie2019/clover/global/consts"
	"github.com/zhanjunjie2019/clover/global/errs"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/app/cmd"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/domain/biserrs"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/domain/model"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/infr/configs"
	"time"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type SadminApp struct {
}

func (s *SadminApp) SadminTokenCreate(ctx context.Context, c cmd.SadminTokenCreateCmd) (rs cmd.SadminTokenCreateResult, err error) {
	auperAdmin := configs.GetAuthConfig().SuperAdmin
	tenant := model.NewTenant(0, model.TenantValue{
		SecretKey:            auperAdmin.SecretKey,
		AccessTokenTimeLimit: auperAdmin.AccessTokenTimeLimit,
	})
	if !tenant.VerifySecretKey(c.SecretKey) {
		err = biserrs.TenantValidationFailedErr
		return
	}
	rs.AccessTokenExpirationTime = time.Now().Unix() + int64(tenant.FullValue().AccessTokenTimeLimit)
	rs.AccessToken, err = uctx.NewJwtClaimsToken("", 0, "", []string{consts.SAdminAuth}, rs.AccessTokenExpirationTime)
	if err != nil {
		err = errs.ToUnifiedError(err)
		return
	}
	return
}
