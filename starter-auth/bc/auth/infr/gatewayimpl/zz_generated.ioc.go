//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package gatewayimpl

import (
	contextx "context"
	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	singleton "github.com/alibaba/ioc-golang/autowire/singleton"
	util "github.com/alibaba/ioc-golang/autowire/util"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/domain/model"
)

func init() {
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &permissionGateway_{}
		},
	})
	permissionGatewayStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &PermissionGateway{}
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
	}
	singleton.RegisterStructDescriptor(permissionGatewayStructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &roleGateway_{}
		},
	})
	roleGatewayStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &RoleGateway{}
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
	}
	singleton.RegisterStructDescriptor(roleGatewayStructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &tenantGateway_{}
		},
	})
	tenantGatewayStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &TenantGateway{}
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
	}
	singleton.RegisterStructDescriptor(tenantGatewayStructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &userGateway_{}
		},
	})
	userGatewayStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &UserGateway{}
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
	}
	singleton.RegisterStructDescriptor(userGatewayStructDescriptor)
}

type permissionGateway_ struct {
	FindByAuthCode_  func(ctx contextx.Context, authCode string) (permission model.Permission, exist bool, err error)
	ListByAuthCodes_ func(ctx contextx.Context, authCodes []string) ([]model.Permission, error)
	Save_            func(ctx contextx.Context, permission model.Permission) (defs.ID, error)
}

func (p *permissionGateway_) FindByAuthCode(ctx contextx.Context, authCode string) (permission model.Permission, exist bool, err error) {
	return p.FindByAuthCode_(ctx, authCode)
}

func (p *permissionGateway_) ListByAuthCodes(ctx contextx.Context, authCodes []string) ([]model.Permission, error) {
	return p.ListByAuthCodes_(ctx, authCodes)
}

func (p *permissionGateway_) Save(ctx contextx.Context, permission model.Permission) (defs.ID, error) {
	return p.Save_(ctx, permission)
}

type roleGateway_ struct {
	SaveSingle_         func(ctx contextx.Context, role model.Role) (defs.ID, error)
	SaveWithPermission_ func(ctx contextx.Context, role model.Role) (id defs.ID, err error)
	FindByID_           func(ctx contextx.Context, id defs.ID) (role model.Role, exist bool, err error)
	FindByRoleCode_     func(ctx contextx.Context, roleCode string) (role model.Role, exist bool, err error)
	ListByRoleCodes_    func(ctx contextx.Context, roleCodes []string) ([]model.Role, error)
}

func (r *roleGateway_) SaveSingle(ctx contextx.Context, role model.Role) (defs.ID, error) {
	return r.SaveSingle_(ctx, role)
}

func (r *roleGateway_) SaveWithPermission(ctx contextx.Context, role model.Role) (id defs.ID, err error) {
	return r.SaveWithPermission_(ctx, role)
}

func (r *roleGateway_) FindByID(ctx contextx.Context, id defs.ID) (role model.Role, exist bool, err error) {
	return r.FindByID_(ctx, id)
}

func (r *roleGateway_) FindByRoleCode(ctx contextx.Context, roleCode string) (role model.Role, exist bool, err error) {
	return r.FindByRoleCode_(ctx, roleCode)
}

func (r *roleGateway_) ListByRoleCodes(ctx contextx.Context, roleCodes []string) ([]model.Role, error) {
	return r.ListByRoleCodes_(ctx, roleCodes)
}

type tenantGateway_ struct {
	FindByTenantID_            func(ctx contextx.Context, tenantID string) (tenant model.Tenant, exist bool, err error)
	ListByTenantIDs_           func(ctx contextx.Context, tenantIDs []string) ([]model.Tenant, error)
	Save_                      func(ctx contextx.Context, tenant model.Tenant) (defs.ID, error)
	PublishInitEvent_          func(ctx contextx.Context, tenant model.Tenant) error
	TenantTablesManualMigrate_ func(ctx contextx.Context) (err error)
}

func (t *tenantGateway_) FindByTenantID(ctx contextx.Context, tenantID string) (tenant model.Tenant, exist bool, err error) {
	return t.FindByTenantID_(ctx, tenantID)
}

func (t *tenantGateway_) ListByTenantIDs(ctx contextx.Context, tenantIDs []string) ([]model.Tenant, error) {
	return t.ListByTenantIDs_(ctx, tenantIDs)
}

func (t *tenantGateway_) Save(ctx contextx.Context, tenant model.Tenant) (defs.ID, error) {
	return t.Save_(ctx, tenant)
}

func (t *tenantGateway_) PublishInitEvent(ctx contextx.Context, tenant model.Tenant) error {
	return t.PublishInitEvent_(ctx, tenant)
}

func (t *tenantGateway_) TenantTablesManualMigrate(ctx contextx.Context) (err error) {
	return t.TenantTablesManualMigrate_(ctx)
}

type userGateway_ struct {
	SaveSingle_                    func(ctx contextx.Context, user model.User) (defs.ID, error)
	SaveWithRole_                  func(ctx contextx.Context, user model.User) (defs.ID, error)
	FindByID_                      func(ctx contextx.Context, id defs.ID) (user model.User, exist bool, err error)
	FindByUserName_                func(ctx contextx.Context, userName string) (user model.User, exist bool, err error)
	ListByByUserNames_             func(ctx contextx.Context, userNames []string) (users []model.User, err error)
	SaveAuthorizationCodeToCache_  func(ctx contextx.Context, user model.User) (authcode string, err error)
	FindByAuthcode_                func(ctx contextx.Context, authcode string) (user model.User, exist bool, err error)
	ListRoleValsByUserID_          func(ctx contextx.Context, userID defs.ID) (roleVals []model.RoleValue, err error)
	ListPermissionValsByRoleCodes_ func(ctx contextx.Context, roleCodes []string) (permissionVals []model.PermissionValue, err error)
}

func (u *userGateway_) SaveSingle(ctx contextx.Context, user model.User) (defs.ID, error) {
	return u.SaveSingle_(ctx, user)
}

func (u *userGateway_) SaveWithRole(ctx contextx.Context, user model.User) (defs.ID, error) {
	return u.SaveWithRole_(ctx, user)
}

func (u *userGateway_) FindByID(ctx contextx.Context, id defs.ID) (user model.User, exist bool, err error) {
	return u.FindByID_(ctx, id)
}

func (u *userGateway_) FindByUserName(ctx contextx.Context, userName string) (user model.User, exist bool, err error) {
	return u.FindByUserName_(ctx, userName)
}

func (u *userGateway_) ListByByUserNames(ctx contextx.Context, userNames []string) (users []model.User, err error) {
	return u.ListByByUserNames_(ctx, userNames)
}

func (u *userGateway_) SaveAuthorizationCodeToCache(ctx contextx.Context, user model.User) (authcode string, err error) {
	return u.SaveAuthorizationCodeToCache_(ctx, user)
}

func (u *userGateway_) FindByAuthcode(ctx contextx.Context, authcode string) (user model.User, exist bool, err error) {
	return u.FindByAuthcode_(ctx, authcode)
}

func (u *userGateway_) ListRoleValsByUserID(ctx contextx.Context, userID defs.ID) (roleVals []model.RoleValue, err error) {
	return u.ListRoleValsByUserID_(ctx, userID)
}

func (u *userGateway_) ListPermissionValsByRoleCodes(ctx contextx.Context, roleCodes []string) (permissionVals []model.PermissionValue, err error) {
	return u.ListPermissionValsByRoleCodes_(ctx, roleCodes)
}

type PermissionGatewayIOCInterface interface {
	FindByAuthCode(ctx contextx.Context, authCode string) (permission model.Permission, exist bool, err error)
	ListByAuthCodes(ctx contextx.Context, authCodes []string) ([]model.Permission, error)
	Save(ctx contextx.Context, permission model.Permission) (defs.ID, error)
}

type RoleGatewayIOCInterface interface {
	SaveSingle(ctx contextx.Context, role model.Role) (defs.ID, error)
	SaveWithPermission(ctx contextx.Context, role model.Role) (id defs.ID, err error)
	FindByID(ctx contextx.Context, id defs.ID) (role model.Role, exist bool, err error)
	FindByRoleCode(ctx contextx.Context, roleCode string) (role model.Role, exist bool, err error)
	ListByRoleCodes(ctx contextx.Context, roleCodes []string) ([]model.Role, error)
}

type TenantGatewayIOCInterface interface {
	FindByTenantID(ctx contextx.Context, tenantID string) (tenant model.Tenant, exist bool, err error)
	ListByTenantIDs(ctx contextx.Context, tenantIDs []string) ([]model.Tenant, error)
	Save(ctx contextx.Context, tenant model.Tenant) (defs.ID, error)
	PublishInitEvent(ctx contextx.Context, tenant model.Tenant) error
	TenantTablesManualMigrate(ctx contextx.Context) (err error)
}

type UserGatewayIOCInterface interface {
	SaveSingle(ctx contextx.Context, user model.User) (defs.ID, error)
	SaveWithRole(ctx contextx.Context, user model.User) (defs.ID, error)
	FindByID(ctx contextx.Context, id defs.ID) (user model.User, exist bool, err error)
	FindByUserName(ctx contextx.Context, userName string) (user model.User, exist bool, err error)
	ListByByUserNames(ctx contextx.Context, userNames []string) (users []model.User, err error)
	SaveAuthorizationCodeToCache(ctx contextx.Context, user model.User) (authcode string, err error)
	FindByAuthcode(ctx contextx.Context, authcode string) (user model.User, exist bool, err error)
	ListRoleValsByUserID(ctx contextx.Context, userID defs.ID) (roleVals []model.RoleValue, err error)
	ListPermissionValsByRoleCodes(ctx contextx.Context, roleCodes []string) (permissionVals []model.PermissionValue, err error)
}

var _permissionGatewaySDID string

func GetPermissionGatewaySingleton() (*PermissionGateway, error) {
	if _permissionGatewaySDID == "" {
		_permissionGatewaySDID = util.GetSDIDByStructPtr(new(PermissionGateway))
	}
	i, err := singleton.GetImpl(_permissionGatewaySDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*PermissionGateway)
	return impl, nil
}

func GetPermissionGatewayIOCInterfaceSingleton() (PermissionGatewayIOCInterface, error) {
	if _permissionGatewaySDID == "" {
		_permissionGatewaySDID = util.GetSDIDByStructPtr(new(PermissionGateway))
	}
	i, err := singleton.GetImplWithProxy(_permissionGatewaySDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(PermissionGatewayIOCInterface)
	return impl, nil
}

type ThisPermissionGateway struct {
}

func (t *ThisPermissionGateway) This() PermissionGatewayIOCInterface {
	thisPtr, _ := GetPermissionGatewayIOCInterfaceSingleton()
	return thisPtr
}

var _roleGatewaySDID string

func GetRoleGatewaySingleton() (*RoleGateway, error) {
	if _roleGatewaySDID == "" {
		_roleGatewaySDID = util.GetSDIDByStructPtr(new(RoleGateway))
	}
	i, err := singleton.GetImpl(_roleGatewaySDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*RoleGateway)
	return impl, nil
}

func GetRoleGatewayIOCInterfaceSingleton() (RoleGatewayIOCInterface, error) {
	if _roleGatewaySDID == "" {
		_roleGatewaySDID = util.GetSDIDByStructPtr(new(RoleGateway))
	}
	i, err := singleton.GetImplWithProxy(_roleGatewaySDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(RoleGatewayIOCInterface)
	return impl, nil
}

type ThisRoleGateway struct {
}

func (t *ThisRoleGateway) This() RoleGatewayIOCInterface {
	thisPtr, _ := GetRoleGatewayIOCInterfaceSingleton()
	return thisPtr
}

var _tenantGatewaySDID string

func GetTenantGatewaySingleton() (*TenantGateway, error) {
	if _tenantGatewaySDID == "" {
		_tenantGatewaySDID = util.GetSDIDByStructPtr(new(TenantGateway))
	}
	i, err := singleton.GetImpl(_tenantGatewaySDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*TenantGateway)
	return impl, nil
}

func GetTenantGatewayIOCInterfaceSingleton() (TenantGatewayIOCInterface, error) {
	if _tenantGatewaySDID == "" {
		_tenantGatewaySDID = util.GetSDIDByStructPtr(new(TenantGateway))
	}
	i, err := singleton.GetImplWithProxy(_tenantGatewaySDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(TenantGatewayIOCInterface)
	return impl, nil
}

type ThisTenantGateway struct {
}

func (t *ThisTenantGateway) This() TenantGatewayIOCInterface {
	thisPtr, _ := GetTenantGatewayIOCInterfaceSingleton()
	return thisPtr
}

var _userGatewaySDID string

func GetUserGatewaySingleton() (*UserGateway, error) {
	if _userGatewaySDID == "" {
		_userGatewaySDID = util.GetSDIDByStructPtr(new(UserGateway))
	}
	i, err := singleton.GetImpl(_userGatewaySDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*UserGateway)
	return impl, nil
}

func GetUserGatewayIOCInterfaceSingleton() (UserGatewayIOCInterface, error) {
	if _userGatewaySDID == "" {
		_userGatewaySDID = util.GetSDIDByStructPtr(new(UserGateway))
	}
	i, err := singleton.GetImplWithProxy(_userGatewaySDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(UserGatewayIOCInterface)
	return impl, nil
}

type ThisUserGateway struct {
}

func (t *ThisUserGateway) This() UserGatewayIOCInterface {
	thisPtr, _ := GetUserGatewayIOCInterfaceSingleton()
	return thisPtr
}