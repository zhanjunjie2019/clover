package repo

import (
	"context"
	"github.com/zhanjunjie2019/clover/global/consts"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/configs"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/repo/po"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
// +ioc:autowire:constructFunc=InitRoleRepo
// +ioc:autowire:implements=github.com/zhanjunjie2019/clover/global/defs.IRepo

type RoleRepo struct {
	TablePre string
}

func InitRoleRepo(r *RoleRepo) (*RoleRepo, error) {
	r.TablePre = "roles"
	return r, nil
}

func (r *RoleRepo) AutoMigrate(ctx context.Context) error {
	tenantID := uctx.GetTenantID(ctx)
	if len(tenantID) == 0 {
		auperAdmin := configs.GetAuthConfig().SuperAdmin
		ctx = context.WithValue(ctx, consts.CtxTenantIDVar, auperAdmin.TenantID)
	}
	return uctx.GetTenantTableDBWithCtx(ctx, r.TablePre).AutoMigrate(po.Role{})
}
