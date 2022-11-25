package gateway

import (
	"context"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/domain/model"
)

type IUserGateway interface {
	Save(ctx context.Context, user model.User) (defs.ID, error)
	FindByUserName(ctx context.Context, userName string) (user model.User, exist bool, err error)
	SaveToCacheByAuthorizationCode(ctx context.Context, user model.User) (authorizationCode string, err error)
}
