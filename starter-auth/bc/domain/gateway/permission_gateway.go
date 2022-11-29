package gateway

import (
	"context"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/domain/model"
)

type IPermissionGateway interface {
	FindByAuthCode(ctx context.Context, authCode string) (permission model.Permission, exist bool, err error)
	Save(ctx context.Context, permission model.Permission) (defs.ID, error)
}
