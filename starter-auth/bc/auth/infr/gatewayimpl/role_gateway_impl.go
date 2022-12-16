package gatewayimpl

import (
	"context"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/utils"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/domain/model"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/infr/gatewayimpl/convs"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/infr/repo"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/infr/repo/po"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type RoleGateway struct {
	RoleRepo              repo.RoleRepoIOCInterface              `singleton:""`
	RolePermissionRelRepo repo.RolePermissionRelRepoIOCInterface `singleton:""`
}

func (r *RoleGateway) SaveSingle(ctx context.Context, role model.Role) (defs.ID, error) {
	return r.RoleRepo.Save(ctx, convs.RoleDOToPO(role))
}

func (r *RoleGateway) SaveWithPermission(ctx context.Context, role model.Role) (id defs.ID, err error) {
	oldrels, err := r.RolePermissionRelRepo.ListByRoleID(ctx, role.ID())
	if err != nil {
		return 0, err
	}
	newrels := convs.RolePermissionDOToPOs(role)
	inserts, updates, deletes := utils.LoadChangeByArrays(newrels, oldrels, func(newObject, oldObject *po.RolePermissionRel) bool {
		if newObject.AuthCode == oldObject.AuthCode {
			newObject.ID = oldObject.ID
			return true
		}
		return false
	})
	if len(inserts) > 0 {
		err = r.RolePermissionRelRepo.BatchInsert(ctx, inserts)
		if err != nil {
			return 0, err
		}
	}
	if len(updates) > 0 {
		err = r.RolePermissionRelRepo.BatchUpdate(ctx, updates)
		if err != nil {
			return 0, err
		}
	}
	if len(deletes) > 0 {
		err = r.RolePermissionRelRepo.BatchDelete(ctx, deletes)
		if err != nil {
			return 0, err
		}
	}
	return r.RoleRepo.Save(ctx, convs.RoleDOToPO(role))
}

func (r *RoleGateway) FindByID(ctx context.Context, id defs.ID) (role model.Role, exist bool, err error) {
	rolePO, exist, err := r.RoleRepo.FindByID(ctx, id)
	if err != nil {
		return nil, false, err
	}
	if exist {
		role = convs.RolePOToDO(rolePO)
	}
	return
}

func (r *RoleGateway) FindByRoleCode(ctx context.Context, roleCode string) (role model.Role, exist bool, err error) {
	rolePO, exist, err := r.RoleRepo.FindByRoleCode(ctx, roleCode)
	if err != nil {
		return nil, false, err
	}
	if exist {
		role = convs.RolePOToDO(rolePO)
	}
	return
}

func (r *RoleGateway) ListByRoleCodes(ctx context.Context, roleCodes []string) ([]model.Role, error) {
	rolePOs, err := r.RoleRepo.ListByRoleCodes(ctx, roleCodes)
	return convs.BatchRolePOToDO(rolePOs), err
}
