package model

import (
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
	// VerifyPassword 验证密码
	VerifyPassword(pwd, salt string) bool
	// EncodePassword 密码加密
	EncodePassword(salt string) string
}

type user struct {
	id         defs.ID
	value      UserValue
	roleValues []RoleValue
}

func (u user) ID() defs.ID {
	return u.id
}

func (u user) FullValue() UserValue {
	return u.value
}

func (u user) GetRoleValues() []RoleValue {
	return u.roleValues
}

func (u *user) SetRoleValues(rvs []RoleValue) {
	u.roleValues = rvs
}

func (u user) VerifyPassword(pwd, salt string) bool {
	return u.value.Password == utils.Md5SaltString(pwd, salt)
}

func (u *user) EncodePassword(salt string) string {
	u.value.Password = utils.Md5SaltString(u.value.Password, salt)
	return u.value.Password
}

type UserValue struct {
	// UserName 用户名
	UserName string
	// Password 密码
	Password string
}
