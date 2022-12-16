package repo

import (
	"context"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/infr/repo/po"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
// +ioc:autowire:constructFunc=InitUserRoleRelRepo

type UserRoleRelRepo struct {
	TablePre string
}

func InitUserRoleRelRepo(u *UserRoleRelRepo) (*UserRoleRelRepo, error) {
	u.TablePre = "user_role_rels"
	return u, nil
}

func (u *UserRoleRelRepo) AutoMigrate(ctx context.Context) error {
	return uctx.GetTenantTableDBWithCtx(ctx, u.TablePre).AutoMigrate(po.UserRoleRel{})
}

func (u *UserRoleRelRepo) ListByUserID(ctx context.Context, userID defs.ID) (rels []po.UserRoleRel, err error) {
	err = uctx.GetTenantTableDBWithCtx(ctx, u.TablePre).Where("user_id=?", userID).Find(&rels).Error
	return
}

func (u *UserRoleRelRepo) BatchInsert(ctx context.Context, pos []po.UserRoleRel) (err error) {
	err = uctx.GetTenantTableDBWithCtx(ctx, u.TablePre).CreateInBatches(pos, 128).Error
	return
}

func (u *UserRoleRelRepo) BatchUpdate(ctx context.Context, pos []po.UserRoleRel) (err error) {
	for i := range pos {
		relpo := pos[i]
		err = uctx.GetTenantTableDBWithCtx(ctx, u.TablePre).Save(&relpo).Error
		if err != nil {
			return
		}
	}
	return
}

func (u *UserRoleRelRepo) BatchDelete(ctx context.Context, pos []po.UserRoleRel) (err error) {
	var ids []defs.ID
	for i := range pos {
		relpo := pos[i]
		ids = append(ids, relpo.ID)
	}
	err = uctx.GetTenantTableDBWithCtx(ctx, u.TablePre).Delete(&po.UserRoleRel{}, ids).Error
	return
}
