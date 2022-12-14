package gateway

import (
	"context"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/domain/model"
)

type IUserGateway interface {
	SaveSingle(ctx context.Context, user model.User) (defs.ID, error)
	SaveWithRole(ctx context.Context, user model.User) (defs.ID, error)
	FindByID(ctx context.Context, id defs.ID) (user model.User, exist bool, err error)
	FindByUserName(ctx context.Context, userName string) (user model.User, exist bool, err error)
	SaveAuthorizationCodeToCache(ctx context.Context, user model.User) (authorizationCode string, err error)
}
