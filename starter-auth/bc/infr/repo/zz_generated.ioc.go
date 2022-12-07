//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package repo

import (
	contextx "context"
	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	singleton "github.com/alibaba/ioc-golang/autowire/singleton"
	util "github.com/alibaba/ioc-golang/autowire/util"
	allimpls "github.com/alibaba/ioc-golang/extension/autowire/allimpls"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/repo/po"
)

func init() {
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &permissionRepo_{}
		},
	})
	permissionRepoStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &PermissionRepo{}
		},
		Metadata: map[string]interface{}{
			"aop": map[string]interface{}{},
			"autowire": map[string]interface{}{
				"common": map[string]interface{}{
					"implements": []interface{}{
						new(defs.IRepo),
					},
				},
			},
		},
	}
	singleton.RegisterStructDescriptor(permissionRepoStructDescriptor)
	allimpls.RegisterStructDescriptor(permissionRepoStructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &rolePermissionRelRepo_{}
		},
	})
	rolePermissionRelRepoStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &RolePermissionRelRepo{}
		},
		ConstructFunc: func(i interface{}, _ interface{}) (interface{}, error) {
			impl := i.(*RolePermissionRelRepo)
			var constructFunc RolePermissionRelRepoConstructFunc = InitRolePermissionRelRepo
			return constructFunc(impl)
		},
		Metadata: map[string]interface{}{
			"aop": map[string]interface{}{},
			"autowire": map[string]interface{}{
				"common": map[string]interface{}{
					"implements": []interface{}{
						new(defs.IRepo),
					},
				},
			},
		},
	}
	singleton.RegisterStructDescriptor(rolePermissionRelRepoStructDescriptor)
	allimpls.RegisterStructDescriptor(rolePermissionRelRepoStructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &roleRepo_{}
		},
	})
	roleRepoStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &RoleRepo{}
		},
		ConstructFunc: func(i interface{}, _ interface{}) (interface{}, error) {
			impl := i.(*RoleRepo)
			var constructFunc RoleRepoConstructFunc = InitRoleRepo
			return constructFunc(impl)
		},
		Metadata: map[string]interface{}{
			"aop": map[string]interface{}{},
			"autowire": map[string]interface{}{
				"common": map[string]interface{}{
					"implements": []interface{}{
						new(defs.IRepo),
					},
				},
			},
		},
	}
	singleton.RegisterStructDescriptor(roleRepoStructDescriptor)
	allimpls.RegisterStructDescriptor(roleRepoStructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &tenantRepo_{}
		},
	})
	tenantRepoStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &TenantRepo{}
		},
		Metadata: map[string]interface{}{
			"aop": map[string]interface{}{},
			"autowire": map[string]interface{}{
				"common": map[string]interface{}{
					"implements": []interface{}{
						new(defs.IRepo),
					},
				},
			},
		},
	}
	singleton.RegisterStructDescriptor(tenantRepoStructDescriptor)
	allimpls.RegisterStructDescriptor(tenantRepoStructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &userRepo_{}
		},
	})
	userRepoStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &UserRepo{}
		},
		ConstructFunc: func(i interface{}, _ interface{}) (interface{}, error) {
			impl := i.(*UserRepo)
			var constructFunc UserRepoConstructFunc = InitUserRepo
			return constructFunc(impl)
		},
		Metadata: map[string]interface{}{
			"aop": map[string]interface{}{},
			"autowire": map[string]interface{}{
				"common": map[string]interface{}{
					"implements": []interface{}{
						new(defs.IRepo),
					},
				},
			},
		},
	}
	singleton.RegisterStructDescriptor(userRepoStructDescriptor)
	allimpls.RegisterStructDescriptor(userRepoStructDescriptor)
}

type RolePermissionRelRepoConstructFunc func(impl *RolePermissionRelRepo) (*RolePermissionRelRepo, error)
type RoleRepoConstructFunc func(impl *RoleRepo) (*RoleRepo, error)
type UserRepoConstructFunc func(impl *UserRepo) (*UserRepo, error)
type permissionRepo_ struct {
	AutoMigrate_     func(ctx contextx.Context) error
	FindByAuthCode_  func(ctx contextx.Context, authCode string) (permissionPO po.Permission, exist bool, err error)
	Save_            func(ctx contextx.Context, permissionPO po.Permission) (defs.ID, error)
	ListByAuthCodes_ func(ctx contextx.Context, authCodes []string) (permissionPOs []po.Permission, err error)
}

func (p *permissionRepo_) AutoMigrate(ctx contextx.Context) error {
	return p.AutoMigrate_(ctx)
}

func (p *permissionRepo_) FindByAuthCode(ctx contextx.Context, authCode string) (permissionPO po.Permission, exist bool, err error) {
	return p.FindByAuthCode_(ctx, authCode)
}

func (p *permissionRepo_) Save(ctx contextx.Context, permissionPO po.Permission) (defs.ID, error) {
	return p.Save_(ctx, permissionPO)
}

func (p *permissionRepo_) ListByAuthCodes(ctx contextx.Context, authCodes []string) (permissionPOs []po.Permission, err error) {
	return p.ListByAuthCodes_(ctx, authCodes)
}

type rolePermissionRelRepo_ struct {
	AutoMigrate_  func(ctx contextx.Context) error
	ListByRoleID_ func(ctx contextx.Context, roleID defs.ID) (rels []po.RolePermissionRel, err error)
	BatchInsert_  func(ctx contextx.Context, pos []po.RolePermissionRel) (err error)
	BatchUpdate_  func(ctx contextx.Context, pos []po.RolePermissionRel) (err error)
	BatchDelete_  func(ctx contextx.Context, pos []po.RolePermissionRel) (err error)
}

func (r *rolePermissionRelRepo_) AutoMigrate(ctx contextx.Context) error {
	return r.AutoMigrate_(ctx)
}

func (r *rolePermissionRelRepo_) ListByRoleID(ctx contextx.Context, roleID defs.ID) (rels []po.RolePermissionRel, err error) {
	return r.ListByRoleID_(ctx, roleID)
}

func (r *rolePermissionRelRepo_) BatchInsert(ctx contextx.Context, pos []po.RolePermissionRel) (err error) {
	return r.BatchInsert_(ctx, pos)
}

func (r *rolePermissionRelRepo_) BatchUpdate(ctx contextx.Context, pos []po.RolePermissionRel) (err error) {
	return r.BatchUpdate_(ctx, pos)
}

func (r *rolePermissionRelRepo_) BatchDelete(ctx contextx.Context, pos []po.RolePermissionRel) (err error) {
	return r.BatchDelete_(ctx, pos)
}

type roleRepo_ struct {
	AutoMigrate_    func(ctx contextx.Context) error
	Save_           func(ctx contextx.Context, rolePO po.Role) (defs.ID, error)
	FindByRoleCode_ func(ctx contextx.Context, roleCode string) (rolePO po.Role, exist bool, err error)
}

func (r *roleRepo_) AutoMigrate(ctx contextx.Context) error {
	return r.AutoMigrate_(ctx)
}

func (r *roleRepo_) Save(ctx contextx.Context, rolePO po.Role) (defs.ID, error) {
	return r.Save_(ctx, rolePO)
}

func (r *roleRepo_) FindByRoleCode(ctx contextx.Context, roleCode string) (rolePO po.Role, exist bool, err error) {
	return r.FindByRoleCode_(ctx, roleCode)
}

type tenantRepo_ struct {
	AutoMigrate_    func(ctx contextx.Context) error
	FindByTenantID_ func(ctx contextx.Context, tenantID string) (tenantPO po.Tenant, exist bool, err error)
	Save_           func(ctx contextx.Context, tenantPO po.Tenant) (defs.ID, error)
}

func (t *tenantRepo_) AutoMigrate(ctx contextx.Context) error {
	return t.AutoMigrate_(ctx)
}

func (t *tenantRepo_) FindByTenantID(ctx contextx.Context, tenantID string) (tenantPO po.Tenant, exist bool, err error) {
	return t.FindByTenantID_(ctx, tenantID)
}

func (t *tenantRepo_) Save(ctx contextx.Context, tenantPO po.Tenant) (defs.ID, error) {
	return t.Save_(ctx, tenantPO)
}

type userRepo_ struct {
	AutoMigrate_    func(ctx contextx.Context) error
	Save_           func(ctx contextx.Context, userPO po.User) (defs.ID, error)
	FindByUserName_ func(ctx contextx.Context, userName string) (userPO po.User, exist bool, err error)
}

func (u *userRepo_) AutoMigrate(ctx contextx.Context) error {
	return u.AutoMigrate_(ctx)
}

func (u *userRepo_) Save(ctx contextx.Context, userPO po.User) (defs.ID, error) {
	return u.Save_(ctx, userPO)
}

func (u *userRepo_) FindByUserName(ctx contextx.Context, userName string) (userPO po.User, exist bool, err error) {
	return u.FindByUserName_(ctx, userName)
}

type PermissionRepoIOCInterface interface {
	AutoMigrate(ctx contextx.Context) error
	FindByAuthCode(ctx contextx.Context, authCode string) (permissionPO po.Permission, exist bool, err error)
	Save(ctx contextx.Context, permissionPO po.Permission) (defs.ID, error)
	ListByAuthCodes(ctx contextx.Context, authCodes []string) (permissionPOs []po.Permission, err error)
}

type RolePermissionRelRepoIOCInterface interface {
	AutoMigrate(ctx contextx.Context) error
	ListByRoleID(ctx contextx.Context, roleID defs.ID) (rels []po.RolePermissionRel, err error)
	BatchInsert(ctx contextx.Context, pos []po.RolePermissionRel) (err error)
	BatchUpdate(ctx contextx.Context, pos []po.RolePermissionRel) (err error)
	BatchDelete(ctx contextx.Context, pos []po.RolePermissionRel) (err error)
}

type RoleRepoIOCInterface interface {
	AutoMigrate(ctx contextx.Context) error
	Save(ctx contextx.Context, rolePO po.Role) (defs.ID, error)
	FindByRoleCode(ctx contextx.Context, roleCode string) (rolePO po.Role, exist bool, err error)
}

type TenantRepoIOCInterface interface {
	AutoMigrate(ctx contextx.Context) error
	FindByTenantID(ctx contextx.Context, tenantID string) (tenantPO po.Tenant, exist bool, err error)
	Save(ctx contextx.Context, tenantPO po.Tenant) (defs.ID, error)
}

type UserRepoIOCInterface interface {
	AutoMigrate(ctx contextx.Context) error
	Save(ctx contextx.Context, userPO po.User) (defs.ID, error)
	FindByUserName(ctx contextx.Context, userName string) (userPO po.User, exist bool, err error)
}

var _permissionRepoSDID string

func GetPermissionRepoSingleton() (*PermissionRepo, error) {
	if _permissionRepoSDID == "" {
		_permissionRepoSDID = util.GetSDIDByStructPtr(new(PermissionRepo))
	}
	i, err := singleton.GetImpl(_permissionRepoSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*PermissionRepo)
	return impl, nil
}

func GetPermissionRepoIOCInterfaceSingleton() (PermissionRepoIOCInterface, error) {
	if _permissionRepoSDID == "" {
		_permissionRepoSDID = util.GetSDIDByStructPtr(new(PermissionRepo))
	}
	i, err := singleton.GetImplWithProxy(_permissionRepoSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(PermissionRepoIOCInterface)
	return impl, nil
}

type ThisPermissionRepo struct {
}

func (t *ThisPermissionRepo) This() PermissionRepoIOCInterface {
	thisPtr, _ := GetPermissionRepoIOCInterfaceSingleton()
	return thisPtr
}

var _rolePermissionRelRepoSDID string

func GetRolePermissionRelRepoSingleton() (*RolePermissionRelRepo, error) {
	if _rolePermissionRelRepoSDID == "" {
		_rolePermissionRelRepoSDID = util.GetSDIDByStructPtr(new(RolePermissionRelRepo))
	}
	i, err := singleton.GetImpl(_rolePermissionRelRepoSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*RolePermissionRelRepo)
	return impl, nil
}

func GetRolePermissionRelRepoIOCInterfaceSingleton() (RolePermissionRelRepoIOCInterface, error) {
	if _rolePermissionRelRepoSDID == "" {
		_rolePermissionRelRepoSDID = util.GetSDIDByStructPtr(new(RolePermissionRelRepo))
	}
	i, err := singleton.GetImplWithProxy(_rolePermissionRelRepoSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(RolePermissionRelRepoIOCInterface)
	return impl, nil
}

type ThisRolePermissionRelRepo struct {
}

func (t *ThisRolePermissionRelRepo) This() RolePermissionRelRepoIOCInterface {
	thisPtr, _ := GetRolePermissionRelRepoIOCInterfaceSingleton()
	return thisPtr
}

var _roleRepoSDID string

func GetRoleRepoSingleton() (*RoleRepo, error) {
	if _roleRepoSDID == "" {
		_roleRepoSDID = util.GetSDIDByStructPtr(new(RoleRepo))
	}
	i, err := singleton.GetImpl(_roleRepoSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*RoleRepo)
	return impl, nil
}

func GetRoleRepoIOCInterfaceSingleton() (RoleRepoIOCInterface, error) {
	if _roleRepoSDID == "" {
		_roleRepoSDID = util.GetSDIDByStructPtr(new(RoleRepo))
	}
	i, err := singleton.GetImplWithProxy(_roleRepoSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(RoleRepoIOCInterface)
	return impl, nil
}

type ThisRoleRepo struct {
}

func (t *ThisRoleRepo) This() RoleRepoIOCInterface {
	thisPtr, _ := GetRoleRepoIOCInterfaceSingleton()
	return thisPtr
}

var _tenantRepoSDID string

func GetTenantRepoSingleton() (*TenantRepo, error) {
	if _tenantRepoSDID == "" {
		_tenantRepoSDID = util.GetSDIDByStructPtr(new(TenantRepo))
	}
	i, err := singleton.GetImpl(_tenantRepoSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*TenantRepo)
	return impl, nil
}

func GetTenantRepoIOCInterfaceSingleton() (TenantRepoIOCInterface, error) {
	if _tenantRepoSDID == "" {
		_tenantRepoSDID = util.GetSDIDByStructPtr(new(TenantRepo))
	}
	i, err := singleton.GetImplWithProxy(_tenantRepoSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(TenantRepoIOCInterface)
	return impl, nil
}

type ThisTenantRepo struct {
}

func (t *ThisTenantRepo) This() TenantRepoIOCInterface {
	thisPtr, _ := GetTenantRepoIOCInterfaceSingleton()
	return thisPtr
}

var _userRepoSDID string

func GetUserRepoSingleton() (*UserRepo, error) {
	if _userRepoSDID == "" {
		_userRepoSDID = util.GetSDIDByStructPtr(new(UserRepo))
	}
	i, err := singleton.GetImpl(_userRepoSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*UserRepo)
	return impl, nil
}

func GetUserRepoIOCInterfaceSingleton() (UserRepoIOCInterface, error) {
	if _userRepoSDID == "" {
		_userRepoSDID = util.GetSDIDByStructPtr(new(UserRepo))
	}
	i, err := singleton.GetImplWithProxy(_userRepoSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(UserRepoIOCInterface)
	return impl, nil
}

type ThisUserRepo struct {
}

func (t *ThisUserRepo) This() UserRepoIOCInterface {
	thisPtr, _ := GetUserRepoIOCInterfaceSingleton()
	return thisPtr
}
