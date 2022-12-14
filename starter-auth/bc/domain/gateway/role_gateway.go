package gateway

import (
	"context"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/domain/model"
)

type IRoleGateway interface {
	SaveSingle(ctx context.Context, role model.Role) (defs.ID, error)
	SaveWithPermission(ctx context.Context, role model.Role) (id defs.ID, err error)
	FindByID(ctx context.Context, id defs.ID) (role model.Role, exist bool, err error)
	FindByRoleCode(ctx context.Context, roleCode string) (role model.Role, exist bool, err error)
	ListByRoleCodes(ctx context.Context, roleCodes []string) (roles []model.Role, err error)
}
