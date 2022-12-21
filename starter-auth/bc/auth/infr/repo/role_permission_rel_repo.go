package repo

import (
	"context"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/infr/repo/po"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
// +ioc:autowire:constructFunc=InitRolePermissionRelRepo

type RolePermissionRelRepo struct {
	TablePre string
}

func InitRolePermissionRelRepo(r *RolePermissionRelRepo) (*RolePermissionRelRepo, error) {
	r.TablePre = "role_permission_rels"
	return r, nil
}

func (r *RolePermissionRelRepo) AutoMigrate(ctx context.Context) error {
	return uctx.GetTenantTableDBWithCtx(ctx, r.TablePre).AutoMigrate(po.RolePermissionRel{})
}

func (r *RolePermissionRelRepo) ListByRoleID(ctx context.Context, roleID defs.ID) (rels []po.RolePermissionRel, err error) {
	err = uctx.GetTenantTableDBWithCtx(ctx, r.TablePre).Where("role_id=?", roleID).Find(&rels).Error
	return
}

func (r *RolePermissionRelRepo) ListByRoleCode(ctx context.Context, roleCode string) (rels []po.RolePermissionRel, err error) {
	err = uctx.GetTenantTableDBWithCtx(ctx, r.TablePre).Where("role_code=?", roleCode).Find(&rels).Error
	return
}

func (r *RolePermissionRelRepo) ListByRoleCodes(ctx context.Context, roleCodes []string) (rels []po.RolePermissionRel, err error) {
	err = uctx.GetTenantTableDBWithCtx(ctx, r.TablePre).Where("role_code IN ?", roleCodes).Find(&rels).Error
	return
}

func (r *RolePermissionRelRepo) BatchInsert(ctx context.Context, pos []po.RolePermissionRel) (err error) {
	err = uctx.GetTenantTableDBWithCtx(ctx, r.TablePre).CreateInBatches(pos, 128).Error
	return
}

func (r *RolePermissionRelRepo) BatchUpdate(ctx context.Context, pos []po.RolePermissionRel) (err error) {
	for i := range pos {
		relpo := pos[i]
		err = uctx.GetTenantTableDBWithCtx(ctx, r.TablePre).Save(&relpo).Error
		if err != nil {
			return
		}
	}
	return
}

func (r *RolePermissionRelRepo) BatchDelete(ctx context.Context, pos []po.RolePermissionRel) (err error) {
	var ids []defs.ID
	for i := range pos {
		relpo := pos[i]
		ids = append(ids, relpo.ID)
	}
	err = uctx.GetTenantTableDBWithCtx(ctx, r.TablePre).Delete(&po.RolePermissionRel{}, ids).Error
	return
}
