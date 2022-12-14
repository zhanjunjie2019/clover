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
	"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/repo/po"
	"time"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type UserGateway struct {
	UserRepo        repo.UserRepoIOCInterface        `singleton:""`
	UserRoleRelRepo repo.UserRoleRelRepoIOCInterface `singleton:""`
	RedisClient     redisc.RedisClientIOCInterface   `singleton:""`
}

func (u *UserGateway) SaveSingle(ctx context.Context, user model.User) (defs.ID, error) {
	return u.UserRepo.Save(ctx, convs.UserDOToPO(user))
}

func (u *UserGateway) SaveWithRole(ctx context.Context, user model.User) (defs.ID, error) {
	oldrels, err := u.UserRoleRelRepo.ListByUserID(ctx, user.ID())
	if err != nil {
		return 0, err
	}
	newrels := convs.UserRoleDOToPO(user)
	inserts, updates, deletes := utils.LoadChangeByArrays(newrels, oldrels, func(newObject, oldObject *po.UserRoleRel) bool {
		if newObject.RoleCode == oldObject.RoleCode {
			newObject.ID = oldObject.ID
			return true
		}
		return false
	})
	if len(inserts) > 0 {
		err = u.UserRoleRelRepo.BatchInsert(ctx, inserts)
		if err != nil {
			return 0, err
		}
	}
	if len(updates) > 0 {
		err = u.UserRoleRelRepo.BatchUpdate(ctx, updates)
		if err != nil {
			return 0, err
		}
	}
	if len(deletes) > 0 {
		err = u.UserRoleRelRepo.BatchDelete(ctx, deletes)
		if err != nil {
			return 0, err
		}
	}
	return u.UserRepo.Save(ctx, convs.UserDOToPO(user))
}

func (u *UserGateway) FindByID(ctx context.Context, id defs.ID) (user model.User, exist bool, err error) {
	userPO, exist, err := u.UserRepo.FindByID(ctx, id)
	if err != nil {
		return nil, false, err
	}
	if exist {
		user = convs.UserPOToDO(userPO)
	}
	return
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

func (u *UserGateway) SaveAuthorizationCodeToCache(ctx context.Context, user model.User) (authorizationCode string, err error) {
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
