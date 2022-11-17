package gatewayimpl

import (
	"context"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/domain/model"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/gatewayimpl/convs"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/repo"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type UserGateway struct {
	UserRepo repo.UserRepoIOCInterface `singleton:""`
}

func (u *UserGateway) FindByUserName(ctx context.Context, userName string) (user model.User, exist bool, err error) {
	userPO, exist, err := u.UserRepo.FindByUserName(ctx, userName)
	if err != nil {
		return nil, false, err
	}
	if exist {
		user = convs.UserPOToDO(userPO)
	}
	return
}
