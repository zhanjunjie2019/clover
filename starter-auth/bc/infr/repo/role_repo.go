package repo

import (
	"context"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/repo/po"
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

func (r *RoleRepo) ManualMigrate(ctx context.Context) error {
	return uctx.GetTenantTableDBWithCtx(ctx, r.TablePre).AutoMigrate(po.Role{})
}
