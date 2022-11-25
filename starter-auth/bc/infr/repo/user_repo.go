package repo

import (
	"context"
	"github.com/zhanjunjie2019/clover/global/consts"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/configs"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/repo/po"
	"gorm.io/gorm"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
// +ioc:autowire:constructFunc=InitUserRepo
// +ioc:autowire:implements=github.com/zhanjunjie2019/clover/global/defs.IRepo

type UserRepo struct {
	TablePre string
}

func InitUserRepo(u *UserRepo) (*UserRepo, error) {
	u.TablePre = "users"
	return u, nil
}

func (u *UserRepo) AutoMigrate(ctx context.Context) error {
	tenantID := uctx.GetTenantID(ctx)
	if len(tenantID) == 0 {
		auperAdmin := configs.GetAuthConfig().SuperAdmin
		ctx = context.WithValue(ctx, consts.CtxTenantIDVar, auperAdmin.TenantID)
	}
	return uctx.GetTenantTableDBWithCtx(ctx, u.TablePre).AutoMigrate(po.User{})
}

func (u *UserRepo) Save(ctx context.Context, userPO po.User) (defs.ID, error) {
	err := uctx.GetTenantTableDBWithCtx(ctx, u.TablePre).Save(&userPO).Error
	return userPO.ID, err
}

func (u *UserRepo) FindByUserName(ctx context.Context, userName string) (userPO po.User, exist bool, err error) {
	err = uctx.GetTenantTableDBWithCtx(ctx, u.TablePre).Where("user_name=?", userName).First(&userPO).Error
	if err == nil {
		exist = true
	} else if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return
}
