package repo

import (
	"context"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/repo/po"
	"gorm.io/gorm"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
// +ioc:autowire:constructFunc=InitRoleRepo

type RoleRepo struct {
	TablePre string
}

func InitRoleRepo(r *RoleRepo) (*RoleRepo, error) {
	r.TablePre = "roles"
	return r, nil
}

func (r *RoleRepo) AutoMigrate(ctx context.Context) error {
	return uctx.GetTenantTableDBWithCtx(ctx, r.TablePre).AutoMigrate(po.Role{})
}

func (r *RoleRepo) Save(ctx context.Context, rolePO po.Role) (defs.ID, error) {
	err := uctx.GetTenantTableDBWithCtx(ctx, r.TablePre).Save(&rolePO).Error
	return rolePO.ID, err
}

func (r *RoleRepo) FindByID(ctx context.Context, id defs.ID) (rolePO po.Role, exist bool, err error) {
	err = uctx.GetTenantTableDBWithCtx(ctx, r.TablePre).Where("id=?", id).First(&rolePO).Error
	if err == nil {
		exist = true
	} else if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return
}

func (r *RoleRepo) FindByRoleCode(ctx context.Context, roleCode string) (rolePO po.Role, exist bool, err error) {
	err = uctx.GetTenantTableDBWithCtx(ctx, r.TablePre).Where("role_code=?", roleCode).First(&rolePO).Error
	if err == nil {
		exist = true
	} else if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return
}

func (r *RoleRepo) ListByRoleCodes(ctx context.Context, roleCodes []string) (rolePOs []po.Role, err error) {
	err = uctx.GetTenantTableDBWithCtx(ctx, r.TablePre).Where("role_code IN ?", roleCodes).Find(&rolePOs).Error
	return
}
