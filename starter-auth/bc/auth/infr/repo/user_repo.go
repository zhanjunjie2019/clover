package repo

import (
	"context"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/infr/repo/po"
	"gorm.io/gorm"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
// +ioc:autowire:constructFunc=InitUserRepo

type UserRepo struct {
	TablePre string
}

func InitUserRepo(u *UserRepo) (*UserRepo, error) {
	u.TablePre = "users"
	return u, nil
}

func (u *UserRepo) AutoMigrate(ctx context.Context) error {
	return uctx.GetTenantTableDBWithCtx(ctx, u.TablePre).AutoMigrate(po.User{})
}

func (u *UserRepo) Save(ctx context.Context, userPO po.User) (defs.ID, error) {
	err := uctx.GetTenantTableDBWithCtx(ctx, u.TablePre).Save(&userPO).Error
	return userPO.ID, err
}

func (u *UserRepo) FindByID(ctx context.Context, id defs.ID) (userPO po.User, exist bool, err error) {
	err = uctx.GetTenantTableDBWithCtx(ctx, u.TablePre).Where("id=?", id).First(&userPO).Error
	if err == nil {
		exist = true
	} else if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return
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

func (u *UserRepo) ListByByUserNames(ctx context.Context, userNames []string) (userPOs []po.User, err error) {
	err = uctx.GetTenantTableDBWithCtx(ctx, u.TablePre).Where("user_name IN ?", userNames).Find(&userPOs).Error
	return
}
