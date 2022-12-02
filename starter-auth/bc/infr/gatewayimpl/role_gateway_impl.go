package gatewayimpl

import (
	"context"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/domain/model"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/gatewayimpl/convs"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/repo"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type RoleGateway struct {
	RoleRepo repo.RoleRepoIOCInterface `singleton:""`
}

func (r *RoleGateway) Save(ctx context.Context, role model.Role) (defs.ID, error) {
	return r.RoleRepo.Save(ctx, convs.RoleDOToPO(role))
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
