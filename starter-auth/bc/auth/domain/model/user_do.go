package model

import (
	"github.com/samber/lo"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/utils"
)

func NewUser(id defs.ID, value UserValue) User {
	return &user{
		id:    id,
		value: value,
	}
}

type User interface {
	ID() defs.ID
	// FullValue 核心数据
	FullValue() UserValue
	// GetRoleValues 获得角色列表
	GetRoleValues() []RoleValue
	// SetRoleValues 赋值角色列表
	SetRoleValues(rvs []RoleValue)
	// GetPermissionValues 获得用户资源列表
	GetPermissionValues() []PermissionValue
	// SetPermissionValues 赋值用户资源列表
	SetPermissionValues(pvs []PermissionValue)
	// VerifyPassword 验证密码
	VerifyPassword(pwd, salt string) bool
	// EncodePassword 密码加密
	EncodePassword(salt string) string
	// GetRoleCodes 获取角色编码列表
	GetRoleCodes() []string
	// GetAuthCodes 获得资源编码列表
	GetAuthCodes() []string
}

type user struct {
	id    defs.ID
	value UserValue
}

func (u user) ID() defs.ID {
	return u.id
}

func (u user) FullValue() UserValue {
	return u.value
}

func (u user) GetRoleValues() []RoleValue {
	return u.value.RoleValues
}

func (u *user) SetRoleValues(rvs []RoleValue) {
	u.value.RoleValues = rvs
}

func (u user) GetPermissionValues() []PermissionValue {
	return u.value.PermissionValues
}

func (u *user) SetPermissionValues(pvs []PermissionValue) {
	u.value.PermissionValues = pvs
}

func (u user) VerifyPassword(pwd, salt string) bool {
	return u.value.Password == utils.Md5SaltString(pwd, salt)
}

func (u *user) EncodePassword(salt string) string {
	u.value.Password = utils.Md5SaltString(u.value.Password, salt)
	return u.value.Password
}

func (u user) GetRoleCodes() []string {
	return lo.Map(u.value.RoleValues, func(item RoleValue, index int) string {
		return item.RoleCode
	})
}

func (u user) GetAuthCodes() []string {
	return lo.Map(u.value.PermissionValues, func(item PermissionValue, index int) string {
		return item.AuthCode
	})
}

type UserValue struct {
	// UserName 用户名
	UserName string
	// Password 密码
	Password string
	// RoleValues 角色列表
	RoleValues []RoleValue
	// PermissionValues 资源许可列表
	PermissionValues []PermissionValue
}
