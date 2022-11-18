package app

import (
	"context"
	"encoding/json"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/errs"
	"github.com/zhanjunjie2019/clover/global/redisc"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"github.com/zhanjunjie2019/clover/global/utils"
	"github.com/zhanjunjie2019/clover/share/auth/dto"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/domain/biserrs"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/domain/gateway"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/domain/model"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/bcconsts"
	"gorm.io/gorm"
	"time"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
// +ioc:autowire:type=allimpls
// +ioc:autowire:implements=github.com/zhanjunjie2019/clover/global/defs.IAppDef

type UserApp struct {
	UserGateway   gateway.IUserGateway           `singleton:"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/gatewayimpl.UserGateway"`
	TenantGateway gateway.ITenantGateway         `singleton:"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/gatewayimpl.TenantGateway"`
	RedisClient   redisc.RedisClientIOCInterface `singleton:""`
	DB            *gorm.DB
}

func (u *UserApp) SetGormDB(db *gorm.DB) {
	u.DB = db
}

func (u *UserApp) UserCreate(ctx context.Context, userName, password string) (id defs.ID, err error) {
	ctx = uctx.WithValueAppDB(ctx, u.DB)
	tenantID := uctx.GetTenantID(ctx)
	_, exist, err := u.UserGateway.FindByUserName(ctx, userName)
	if err != nil {
		err = errs.ToUnifiedError(err)
		return
	}
	if exist {
		err = biserrs.UserAlreadyExistsErr
		return
	}
	user := model.NewUser(0, model.UserValue{
		UserName: userName,
		Password: password,
	})
	user.EncodePassword(tenantID + "@" + userName)
	id, err = u.UserGateway.Save(ctx, user)
	if err != nil {
		err = errs.ToUnifiedError(err)
		return
	}
	return
}

// UserAuthorizationCode 登录验证用户账号密码，验证通过后在Redis保存一个授权码60秒有效，关联用户信息。用以可以用授权码接口换取登录Token。
func (u *UserApp) UserAuthorizationCode(ctx context.Context, userName, password, redirect string) (authorizationCode, redirectUrl string, err error) {
	ctx = uctx.WithValueAppDB(ctx, u.DB)
	tenantID := uctx.GetTenantID(ctx)
	user, exist, err := u.UserGateway.FindByUserName(ctx, userName)
	if err != nil {
		err = errs.ToUnifiedError(err)
		return
	}
	if !exist {
		err = biserrs.LoginVerifyFailedErr
		return
	}
	verifyPassword := user.VerifyPassword(password, tenantID+"@"+userName)
	if !verifyPassword {
		err = biserrs.LoginVerifyFailedErr
		return
	}
	if len(redirect) == 0 {
		var tenant model.Tenant
		tenant, exist, err = u.TenantGateway.FindByTenantID(ctx, tenantID)
		if err != nil {
			err = errs.ToUnifiedError(err)
			return
		}
		if !exist {
			err = biserrs.LoginVerifyFailedErr
			return
		}
		redirectUrl = tenant.FullValue().RedirectUrl
	} else {
		redirectUrl = redirect
	}
	authorizationCode = utils.UUID()
	userAuthorizationCode := dto.UserAuthorizationCode{
		TenantID: tenantID,
		UserID:   user.ID(),
		UserName: userName,
	}
	client, err := u.RedisClient.GetClient()
	if err != nil {
		err = errs.ToUnifiedError(err)
		return
	}
	bytes, err := json.Marshal(userAuthorizationCode)
	if err != nil {
		err = errs.ToUnifiedError(err)
		return
	}
	err = client.Set(ctx, bcconsts.RedisAuthCodePre+authorizationCode, bytes, time.Minute).Err()
	if err != nil {
		err = errs.ToUnifiedError(err)
		return
	}
	return
}
