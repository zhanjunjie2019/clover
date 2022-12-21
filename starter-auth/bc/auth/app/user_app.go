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
	"time"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
// +ioc:autowire:type=allimpls
// +ioc:autowire:implements=github.com/zhanjunjie2019/clover/global/defs.IAppDef

type UserApp struct {
	UserGateway   gateway.IUserGateway   `singleton:"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/infr/gatewayimpl.UserGateway"`
	RoleGateway   gateway.IRoleGateway   `singleton:"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/infr/gatewayimpl.RoleGateway"`
	TenantGateway gateway.ITenantGateway `singleton:"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/infr/gatewayimpl.TenantGateway"`
	DB            *gorm.DB
}

func (u *UserApp) SetGormDB(db *gorm.DB) {
	u.DB = db
}

func (u *UserApp) UserCreate(ctx context.Context, c cmd.UserCreateCmd) (rs cmd.UserCreateResult, err error) {
	err = u.DB.Transaction(func(tx *gorm.DB) (err error) {
		ctx = uctx.WithValueAppDB(ctx, tx)
		tenantID := uctx.GetTenantID(ctx)
		userNames := lo.Map(c.Users, func(item cmd.UserInfo, index int) string {
			return item.UserName
		})
		users, err := u.UserGateway.ListByByUserNames(ctx, userNames)
		if err != nil {
			err = errs.ToUnifiedError(err)
			return
		}
		if len(users) > 0 {
			userNames = lo.Map(users, func(item model.User, index int) string {
				return item.FullValue().UserName
			})
			err = biserrs.UserAlreadyExistsErrWithUserName(userNames...)
			return
		}
		for _, userInfo := range c.Users {
			user := model.NewUser(0, model.UserValue{
				UserName: userInfo.UserName,
				Password: userInfo.Password,
			})
			user.EncodePassword(tenantID + "@" + userInfo.UserName)
			var id defs.ID
			id, err = u.UserGateway.SaveSingle(ctx, user)
			if err != nil {
				err = errs.ToUnifiedError(err)
				return
			}
			rs.UserIDs = append(rs.UserIDs, id)
		}
		return nil
	})
	return
}

// UserAuthorizationCode 登录验证用户账号密码，验证通过后在Redis保存一个授权码60秒有效，关联用户信息。用以可以用授权码接口换取登录Token。
func (u *UserApp) UserAuthorizationCode(ctx context.Context, c cmd.UserAuthorizationCodeCmd) (rs cmd.UserAuthorizationCodeResult, err error) {
	ctx = uctx.WithValueAppDB(ctx, u.DB)
	tenantID := uctx.GetTenantID(ctx)
	user, exist, err := u.UserGateway.FindByUserName(ctx, c.UserName)
	if err != nil {
		err = errs.ToUnifiedError(err)
		return
	}
	if !exist {
		err = biserrs.LoginVerifyFailedErr
		return
	}
	verifyPassword := user.VerifyPassword(c.Password, tenantID+"@"+c.UserName)
	if !verifyPassword {
		err = biserrs.LoginVerifyFailedErr
		return
	}
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
	rs.RedirectUrl = tenant.FullValue().RedirectUrl
	rs.AuthorizationCode, err = u.UserGateway.SaveAuthorizationCodeToCache(ctx, user)
	if err != nil {
		err = errs.ToUnifiedError(err)
		return
	}
	return
}

func (u *UserApp) UserTokenByAuthcode(ctx context.Context, c cmd.UserTokenByAuthcodeCmd) (rs cmd.UserTokenByAuthcodeResult, err error) {
	ctx = uctx.WithValueAppDB(ctx, u.DB)
	tenantID := uctx.GetTenantID(ctx)
	tenant, exist, err := u.TenantGateway.FindByTenantID(ctx, tenantID)
	if err != nil {
		err = errs.ToUnifiedError(err)
		return
	}
	if !exist {
		err = biserrs.TenantDoesNotExistErrWithTenantID(tenantID)
		return
	}
	user, exist, err := u.UserGateway.FindByAuthcode(ctx, c.AuthorizationCode)
	if err != nil {
		err = errs.ToUnifiedError(err)
		return
	}
	if !exist {
		err = biserrs.LoginVerifyFailedErr
		return
	}
	rs.AccessTokenExpirationTime = time.Now().Add(time.Second * time.Duration(tenant.FullValue().AccessTokenTimeLimit)).Unix()
	rs.AccessToken, err = uctx.NewJwtClaimsToken(tenantID, user.ID().UInt64(), user.FullValue().UserName, user.GetAuthCodes(), rs.AccessTokenExpirationTime)
	if err != nil {
		err = errs.ToUnifiedError(err)
		return
	}
	return
}

func (u *UserApp) UserRoleAssignment(ctx context.Context, c cmd.UserRoleAssignmentCmd) (rs cmd.UserRoleAssignmentResult, err error) {
	err = u.DB.Transaction(func(tx *gorm.DB) (err error) {
		ctx = uctx.WithValueAppDB(ctx, tx)
		// 验证有效性
		var (
			user  model.User
			exist bool
		)
		if c.UserID == 0 {
			user, exist, err = u.UserGateway.FindByUserName(ctx, c.UserName)
		} else {
			user, exist, err = u.UserGateway.FindByID(ctx, c.UserID)
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
		roles, err := u.RoleGateway.ListByRoleCodes(ctx, c.RoleCodes)
		if err != nil {
			err = errs.ToUnifiedError(err)
			return
		}
		var rvs []model.RoleValue
		for _, roleCode := range c.RoleCodes {
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
		// 校验完成后赋值
		user.SetRoleValues(rvs)
		// 保存聚合根
		rs.UserID, err = u.UserGateway.SaveWithRole(ctx, user)
		if err != nil {
			err = errs.ToUnifiedError(err)
			return
		}
		return nil
	})
	return
}
