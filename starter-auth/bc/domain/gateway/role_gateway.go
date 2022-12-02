package gateway

import (
	"context"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/domain/model"
)

type IRoleGateway interface {
	Save(ctx context.Context, role model.Role) (defs.ID, error)
	FindByRoleCode(ctx context.Context, roleCode string) (role model.Role, exist bool, err error)
}
