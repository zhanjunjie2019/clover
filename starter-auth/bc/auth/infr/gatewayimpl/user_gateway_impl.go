package gatewayimpl

import (
	"context"
	"github.com/gogo/protobuf/proto"
	"github.com/samber/lo"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/redisc"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"github.com/zhanjunjie2019/clover/global/utils"
	"github.com/zhanjunjie2019/clover/share/auth/protobuf"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/domain/model"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/infr/bcconsts"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/infr/gatewayimpl/convs"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/infr/repo"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/infr/repo/po"
	"time"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type UserGateway struct {
	UserRepo              repo.UserRepoIOCInterface              `singleton:""`
	UserRoleRelRepo       repo.UserRoleRelRepoIOCInterface       `singleton:""`
	RoleRepo              repo.RoleRepoIOCInterface              `singleton:""`
	RolePermissionRelRepo repo.RolePermissionRelRepoIOCInterface `singleton:""`
	PermissionRepo        repo.PermissionRepoIOCInterface        `singleton:""`
	RedisClient           redisc.RedisClientIOCInterface         `singleton:""`
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

func (u *UserGateway) ListByByUserNames(ctx context.Context, userNames []string) (users []model.User, err error) {
	userPOs, err := u.UserRepo.ListByByUserNames(ctx, userNames)
	return convs.BathUserPOToDO(userPOs), err
}

func (u *UserGateway) SaveAuthorizationCodeToCache(ctx context.Context, user model.User) (authcode string, err error) {
	authcode = utils.TinyUUID()
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
	err = client.Set(ctx, bcconsts.RedisAuthCodePre+authcode, bytes, time.Minute).Err()
	if err != nil {
		return
	}
	return
}

func (u *UserGateway) FindByAuthcode(ctx context.Context, authcode string) (user model.User, exist bool, err error) {
	// 链接redis获取用户缓存信息
	client, err := u.RedisClient.GetClient()
	if err != nil {
		return
	}
	bytes, err := client.Get(ctx, bcconsts.RedisAuthCodePre+authcode).Bytes()
	if err != nil {
		return
	}
	var userAuthcode protobuf.UserAuthorizationCode
	err = proto.Unmarshal(bytes, &userAuthcode)
	if err != nil {
		return
	}
	// 验证租户ID是否一直
	tenantID := uctx.GetTenantID(ctx)
	if tenantID != userAuthcode.TenantID {
		return
	}
	// 查询数据库用户信息
	userPO, exist, err := u.UserRepo.FindByID(ctx, defs.ID(userAuthcode.UserID))
	if err != nil {
		return
	}
	if !exist {
		return
	}
	user = convs.UserPOToDO(userPO)
	// 查询数据库角色信息
	roleVals, err := u.ListRoleValsByUserID(ctx, user.ID())
	if err != nil {
		return
	}
	user.SetRoleValues(roleVals)
	roleCodes := lo.Map(roleVals, func(item model.RoleValue, index int) string {
		return item.RoleCode
	})
	permissionVals, err := u.ListPermissionValsByRoleCodes(ctx, roleCodes)
	if err != nil {
		return
	}
	user.SetPermissionValues(permissionVals)
	return
}

// ListRoleValsByUserID 根据用户ID查询角色值对象
func (u *UserGateway) ListRoleValsByUserID(ctx context.Context, userID defs.ID) (roleVals []model.RoleValue, err error) {
	roleRels, err := u.UserRoleRelRepo.ListByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	for _, roleRel := range roleRels {
		roleVals = append(roleVals, convs.UserRolePOToValue(roleRel))
	}
	return
}

// ListPermissionValsByRoleCodes 根据角色编码查询资源权限值对象
func (u *UserGateway) ListPermissionValsByRoleCodes(ctx context.Context, roleCodes []string) (permissionVals []model.PermissionValue, err error) {
	permissionRels, err := u.RolePermissionRelRepo.ListByRoleCodes(ctx, roleCodes)
	if err != nil {
		return nil, err
	}
	for _, permissionRel := range permissionRels {
		permissionVals = append(permissionVals, convs.RolePermissionPOToValue(permissionRel))
	}
	return
}
