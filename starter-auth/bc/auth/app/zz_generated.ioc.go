//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package app

import (
	contextx "context"
	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	singleton "github.com/alibaba/ioc-golang/autowire/singleton"
	util "github.com/alibaba/ioc-golang/autowire/util"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/app/cmd"
)

func init() {
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &permissionApp_{}
		},
	})
	permissionAppStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &PermissionApp{}
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
	}
	singleton.RegisterStructDescriptor(permissionAppStructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &roleApp_{}
		},
	})
	roleAppStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &RoleApp{}
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
	}
	singleton.RegisterStructDescriptor(roleAppStructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &sadminApp_{}
		},
	})
	sadminAppStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &SadminApp{}
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
	}
	singleton.RegisterStructDescriptor(sadminAppStructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &tenantApp_{}
		},
	})
	tenantAppStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &TenantApp{}
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
	}
	singleton.RegisterStructDescriptor(tenantAppStructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &userApp_{}
		},
	})
	userAppStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &UserApp{}
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
	}
	singleton.RegisterStructDescriptor(userAppStructDescriptor)
}

type permissionApp_ struct {
	PermissionCreate_ func(ctx contextx.Context, c cmd.PermissionCreateCmd) (rs cmd.PermissionCreateResult, err error)
}

func (p *permissionApp_) PermissionCreate(ctx contextx.Context, c cmd.PermissionCreateCmd) (rs cmd.PermissionCreateResult, err error) {
	return p.PermissionCreate_(ctx, c)
}

type roleApp_ struct {
	RoleCreate_               func(ctx contextx.Context, c cmd.RoleCreateCmd) (rs cmd.RoleCreateResult, err error)
	RolePermissionAssignment_ func(ctx contextx.Context, c cmd.RolePermissionAssignmentCmd) (rs cmd.RolePermissionAssignmentResult, err error)
}

func (r *roleApp_) RoleCreate(ctx contextx.Context, c cmd.RoleCreateCmd) (rs cmd.RoleCreateResult, err error) {
	return r.RoleCreate_(ctx, c)
}

func (r *roleApp_) RolePermissionAssignment(ctx contextx.Context, c cmd.RolePermissionAssignmentCmd) (rs cmd.RolePermissionAssignmentResult, err error) {
	return r.RolePermissionAssignment_(ctx, c)
}

type sadminApp_ struct {
	SadminTokenCreate_ func(ctx contextx.Context, c cmd.SadminTokenCreateCmd) (rs cmd.SadminTokenCreateResult, err error)
}

func (s *sadminApp_) SadminTokenCreate(ctx contextx.Context, c cmd.SadminTokenCreateCmd) (rs cmd.SadminTokenCreateResult, err error) {
	return s.SadminTokenCreate_(ctx, c)
}

type tenantApp_ struct {
	TenantCreate_      func(ctx contextx.Context, c cmd.TenantCreateCmd) (rs cmd.TenantCreateResult, err error)
	TenantInit_        func(ctx contextx.Context, c cmd.TenantInitCmd) (err error)
	TenantTokenCreate_ func(ctx contextx.Context, c cmd.TenantTokenCreateCmd) (rs cmd.TenantTokenCreateResult, err error)
}

func (t *tenantApp_) TenantCreate(ctx contextx.Context, c cmd.TenantCreateCmd) (rs cmd.TenantCreateResult, err error) {
	return t.TenantCreate_(ctx, c)
}

func (t *tenantApp_) TenantInit(ctx contextx.Context, c cmd.TenantInitCmd) (err error) {
	return t.TenantInit_(ctx, c)
}

func (t *tenantApp_) TenantTokenCreate(ctx contextx.Context, c cmd.TenantTokenCreateCmd) (rs cmd.TenantTokenCreateResult, err error) {
	return t.TenantTokenCreate_(ctx, c)
}

type userApp_ struct {
	UserCreate_            func(ctx contextx.Context, c cmd.UserCreateCmd) (rs cmd.UserCreateResult, err error)
	UserAuthorizationCode_ func(ctx contextx.Context, c cmd.UserAuthorizationCodeCmd) (rs cmd.UserAuthorizationCodeResult, err error)
	UserTokenByAuthcode_   func(ctx contextx.Context, c cmd.UserTokenByAuthcodeCmd) (rs cmd.UserTokenByAuthcodeResult, err error)
	UserRoleAssignment_    func(ctx contextx.Context, c cmd.UserRoleAssignmentCmd) (rs cmd.UserRoleAssignmentResult, err error)
}

func (u *userApp_) UserCreate(ctx contextx.Context, c cmd.UserCreateCmd) (rs cmd.UserCreateResult, err error) {
	return u.UserCreate_(ctx, c)
}

func (u *userApp_) UserAuthorizationCode(ctx contextx.Context, c cmd.UserAuthorizationCodeCmd) (rs cmd.UserAuthorizationCodeResult, err error) {
	return u.UserAuthorizationCode_(ctx, c)
}

func (u *userApp_) UserTokenByAuthcode(ctx contextx.Context, c cmd.UserTokenByAuthcodeCmd) (rs cmd.UserTokenByAuthcodeResult, err error) {
	return u.UserTokenByAuthcode_(ctx, c)
}

func (u *userApp_) UserRoleAssignment(ctx contextx.Context, c cmd.UserRoleAssignmentCmd) (rs cmd.UserRoleAssignmentResult, err error) {
	return u.UserRoleAssignment_(ctx, c)
}

type PermissionAppIOCInterface interface {
	PermissionCreate(ctx contextx.Context, c cmd.PermissionCreateCmd) (rs cmd.PermissionCreateResult, err error)
}

type RoleAppIOCInterface interface {
	RoleCreate(ctx contextx.Context, c cmd.RoleCreateCmd) (rs cmd.RoleCreateResult, err error)
	RolePermissionAssignment(ctx contextx.Context, c cmd.RolePermissionAssignmentCmd) (rs cmd.RolePermissionAssignmentResult, err error)
}

type SadminAppIOCInterface interface {
	SadminTokenCreate(ctx contextx.Context, c cmd.SadminTokenCreateCmd) (rs cmd.SadminTokenCreateResult, err error)
}

type TenantAppIOCInterface interface {
	TenantCreate(ctx contextx.Context, c cmd.TenantCreateCmd) (rs cmd.TenantCreateResult, err error)
	TenantInit(ctx contextx.Context, c cmd.TenantInitCmd) (err error)
	TenantTokenCreate(ctx contextx.Context, c cmd.TenantTokenCreateCmd) (rs cmd.TenantTokenCreateResult, err error)
}

type UserAppIOCInterface interface {
	UserCreate(ctx contextx.Context, c cmd.UserCreateCmd) (rs cmd.UserCreateResult, err error)
	UserAuthorizationCode(ctx contextx.Context, c cmd.UserAuthorizationCodeCmd) (rs cmd.UserAuthorizationCodeResult, err error)
	UserTokenByAuthcode(ctx contextx.Context, c cmd.UserTokenByAuthcodeCmd) (rs cmd.UserTokenByAuthcodeResult, err error)
	UserRoleAssignment(ctx contextx.Context, c cmd.UserRoleAssignmentCmd) (rs cmd.UserRoleAssignmentResult, err error)
}

var _permissionAppSDID string

func GetPermissionAppSingleton() (*PermissionApp, error) {
	if _permissionAppSDID == "" {
		_permissionAppSDID = util.GetSDIDByStructPtr(new(PermissionApp))
	}
	i, err := singleton.GetImpl(_permissionAppSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*PermissionApp)
	return impl, nil
}

func GetPermissionAppIOCInterfaceSingleton() (PermissionAppIOCInterface, error) {
	if _permissionAppSDID == "" {
		_permissionAppSDID = util.GetSDIDByStructPtr(new(PermissionApp))
	}
	i, err := singleton.GetImplWithProxy(_permissionAppSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(PermissionAppIOCInterface)
	return impl, nil
}

type ThisPermissionApp struct {
}

func (t *ThisPermissionApp) This() PermissionAppIOCInterface {
	thisPtr, _ := GetPermissionAppIOCInterfaceSingleton()
	return thisPtr
}

var _roleAppSDID string

func GetRoleAppSingleton() (*RoleApp, error) {
	if _roleAppSDID == "" {
		_roleAppSDID = util.GetSDIDByStructPtr(new(RoleApp))
	}
	i, err := singleton.GetImpl(_roleAppSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*RoleApp)
	return impl, nil
}

func GetRoleAppIOCInterfaceSingleton() (RoleAppIOCInterface, error) {
	if _roleAppSDID == "" {
		_roleAppSDID = util.GetSDIDByStructPtr(new(RoleApp))
	}
	i, err := singleton.GetImplWithProxy(_roleAppSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(RoleAppIOCInterface)
	return impl, nil
}

type ThisRoleApp struct {
}

func (t *ThisRoleApp) This() RoleAppIOCInterface {
	thisPtr, _ := GetRoleAppIOCInterfaceSingleton()
	return thisPtr
}

var _sadminAppSDID string

func GetSadminAppSingleton() (*SadminApp, error) {
	if _sadminAppSDID == "" {
		_sadminAppSDID = util.GetSDIDByStructPtr(new(SadminApp))
	}
	i, err := singleton.GetImpl(_sadminAppSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*SadminApp)
	return impl, nil
}

func GetSadminAppIOCInterfaceSingleton() (SadminAppIOCInterface, error) {
	if _sadminAppSDID == "" {
		_sadminAppSDID = util.GetSDIDByStructPtr(new(SadminApp))
	}
	i, err := singleton.GetImplWithProxy(_sadminAppSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(SadminAppIOCInterface)
	return impl, nil
}

type ThisSadminApp struct {
}

func (t *ThisSadminApp) This() SadminAppIOCInterface {
	thisPtr, _ := GetSadminAppIOCInterfaceSingleton()
	return thisPtr
}

var _tenantAppSDID string

func GetTenantAppSingleton() (*TenantApp, error) {
	if _tenantAppSDID == "" {
		_tenantAppSDID = util.GetSDIDByStructPtr(new(TenantApp))
	}
	i, err := singleton.GetImpl(_tenantAppSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*TenantApp)
	return impl, nil
}

func GetTenantAppIOCInterfaceSingleton() (TenantAppIOCInterface, error) {
	if _tenantAppSDID == "" {
		_tenantAppSDID = util.GetSDIDByStructPtr(new(TenantApp))
	}
	i, err := singleton.GetImplWithProxy(_tenantAppSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(TenantAppIOCInterface)
	return impl, nil
}

type ThisTenantApp struct {
}

func (t *ThisTenantApp) This() TenantAppIOCInterface {
	thisPtr, _ := GetTenantAppIOCInterfaceSingleton()
	return thisPtr
}

var _userAppSDID string

func GetUserAppSingleton() (*UserApp, error) {
	if _userAppSDID == "" {
		_userAppSDID = util.GetSDIDByStructPtr(new(UserApp))
	}
	i, err := singleton.GetImpl(_userAppSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*UserApp)
	return impl, nil
}

func GetUserAppIOCInterfaceSingleton() (UserAppIOCInterface, error) {
	if _userAppSDID == "" {
		_userAppSDID = util.GetSDIDByStructPtr(new(UserApp))
	}
	i, err := singleton.GetImplWithProxy(_userAppSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(UserAppIOCInterface)
	return impl, nil
}

type ThisUserApp struct {
}

func (t *ThisUserApp) This() UserAppIOCInterface {
	thisPtr, _ := GetUserAppIOCInterfaceSingleton()
	return thisPtr
}
