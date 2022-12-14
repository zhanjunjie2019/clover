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

type PermissionGateway struct {
	PermissionRepo repo.PermissionRepoIOCInterface `singleton:""`
}

func (p *PermissionGateway) FindByAuthCode(ctx context.Context, authCode string) (permission model.Permission, exist bool, err error) {
	permissionPO, exist, err := p.PermissionRepo.FindByAuthCode(ctx, authCode)
	if err != nil {
		return nil, false, err
	}
	if exist {
		permission = convs.PermissionPOToDO(permissionPO)
	}
	return
}

func (p *PermissionGateway) Save(ctx context.Context, permission model.Permission) (defs.ID, error) {
	return p.PermissionRepo.Save(ctx, convs.PermissionDOToPO(permission))
}

func (p *PermissionGateway) ListByAuthCodes(ctx context.Context, authCodes []string) ([]model.Permission, error) {
	permissionPOs, err := p.PermissionRepo.ListByAuthCodes(ctx, authCodes)
	return convs.BatchPermissionPOToDO(permissionPOs), err
}
