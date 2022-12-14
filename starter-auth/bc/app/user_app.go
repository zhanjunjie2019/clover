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

type UserApp struct {
	UserGateway   gateway.IUserGateway   `singleton:"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/gatewayimpl.UserGateway"`
	RoleGateway   gateway.IRoleGateway   `singleton:"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/gatewayimpl.RoleGateway"`
	TenantGateway gateway.ITenantGateway `singleton:"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/gatewayimpl.TenantGateway"`
	DB            *gorm.DB
}

func (u *UserApp) SetGormDB(db *gorm.DB) {
	u.DB = db
}

func (u *UserApp) UserCreate(ctx context.Context, userName, password string) (id defs.ID, err error) {
	err = u.DB.Transaction(func(tx *gorm.DB) (err error) {
		ctx = uctx.WithValueAppDB(ctx, tx)
		tenantID := uctx.GetTenantID(ctx)
		_, exist, err := u.UserGateway.FindByUserName(ctx, userName)
		if err != nil {
			err = errs.ToUnifiedError(err)
			return
		}
		if exist {
			err = biserrs.UserAlreadyExistsErrWithUserName(userName)
			return
		}
		user := model.NewUser(0, model.UserValue{
			UserName: userName,
			Password: password,
		})
		user.EncodePassword(tenantID + "@" + userName)
		id, err = u.UserGateway.SaveSingle(ctx, user)
		if err != nil {
			err = errs.ToUnifiedError(err)
			return
		}
		return nil
	})
	return
}

// UserAuthorizationCode 登录验证用户账号密码，验证通过后在Redis保存一个授权码60秒有效，关联用户信息。用以可以用授权码接口换取登录Token。
func (u *UserApp) UserAuthorizationCode(ctx context.Context, userName, password, redirect string) (authorizationCode, redirectUrl string, err error) {
	err = u.DB.Transaction(func(tx *gorm.DB) (err error) {
		ctx = uctx.WithValueAppDB(ctx, tx)
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
		authorizationCode, err = u.UserGateway.SaveAuthorizationCodeToCache(ctx, user)
		if err != nil {
			err = errs.ToUnifiedError(err)
			return
		}
		return nil
	})
	return
}

func (u *UserApp) UserRoleAssignment(ctx context.Context, userID defs.ID, userCode string, roleCodes []string) (id defs.ID, err error) {
	err = u.DB.Transaction(func(tx *gorm.DB) (err error) {
		ctx = uctx.WithValueAppDB(ctx, tx)
		// 验证有效性
		var (
			user  model.User
			exist bool
		)
		if userID == 0 {
			user, exist, err = u.UserGateway.FindByUserName(ctx, userCode)
		} else {
			user, exist, err = u.UserGateway.FindByID(ctx, userID)
		}
		if err != nil {
			err = errs.ToUnifiedError(err)
			return
		}
		if !exist {
			err = biserrs.UserDoesNotExistErr
			return
		}
		// 获取角色列表
		roles, err := u.RoleGateway.ListByRoleCodes(ctx, roleCodes)
		if err != nil {
			err = errs.ToUnifiedError(err)
			return
		}
		var rvs []model.RoleValue
		for i := range roleCodes {
			roleCode := roleCodes[i]
			var rv *model.RoleValue
			for ri := range roles {
				val := roles[ri].FullValue()
				if val.RoleCode == roleCode {
					rv = &val
					break
				}
			}
			if rv == nil {
				err = biserrs.RoleDoesNotExistErrWithRoleCode(roleCode)
				return
			}
			rvs = append(rvs, *rv)
		}
		user.SetRoleValues(rvs)
		id, err = u.UserGateway.SaveWithRole(ctx, user)
		if err != nil {
			err = errs.ToUnifiedError(err)
			return
		}
		return nil
	})
	return
}
