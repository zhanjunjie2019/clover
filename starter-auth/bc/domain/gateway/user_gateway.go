package gateway

import (
	"context"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/domain/model"
)

type IUserGateway interface {
	FindByUserName(ctx context.Context, userName string) (user model.User, exist bool, err error)
}
