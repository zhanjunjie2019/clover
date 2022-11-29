package gatewayimpl

import (
	"context"
	"github.com/gogo/protobuf/proto"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/redisc"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"github.com/zhanjunjie2019/clover/global/utils"
	"github.com/zhanjunjie2019/clover/share/auth/protobuf"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/domain/model"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/bcconsts"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/gatewayimpl/convs"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/repo"
	"time"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type UserGateway struct {
	UserRepo    repo.UserRepoIOCInterface      `singleton:""`
	RedisClient redisc.RedisClientIOCInterface `singleton:""`
}

func (u *UserGateway) Save(ctx context.Context, user model.User) (defs.ID, error) {
	return u.UserRepo.Save(ctx, convs.UserDOToPO(user))
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

func (u *UserGateway) SaveToCacheByAuthorizationCode(ctx context.Context, user model.User) (authorizationCode string, err error) {
	authorizationCode = utils.UUID()
	userAuthorizationCode := protobuf.UserAuthorizationCode{
		TenantID: uctx.GetTenantID(ctx),
		UserID:   user.ID().UInt64(),
		UserName: user.FullValue().UserName,
	}
	client, err := u.RedisClient.GetClient()
	if err != nil {
		return
	}
	bytes, err := proto.Marshal(&userAuthorizationCode)
	if err != nil {
		return
	}
	err = client.Set(ctx, bcconsts.RedisAuthCodePre+authorizationCode, bytes, time.Minute).Err()
	if err != nil {
		return
	}
	return
}
