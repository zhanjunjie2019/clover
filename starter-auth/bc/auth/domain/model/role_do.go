package model

import "github.com/zhanjunjie2019/clover/global/defs"

func NewRole(id defs.ID, value RoleValue) Role {
	return &role{
		id:    id,
		value: value,
	}
}

type Role interface {
	ID() defs.ID
	// FullValue 核心数据
	FullValue() RoleValue
	// GetPermissionValues 获得角色资源许可
	GetPermissionValues() []PermissionValue
	// SetPermissionValues 赋值角色资源许可
	SetPermissionValues([]PermissionValue)
}

type role struct {
	id    defs.ID
	value RoleValue
}

func (r role) ID() defs.ID {
	return r.id
}

func (r role) FullValue() RoleValue {
	return r.value
}

func (r role) GetPermissionValues() []PermissionValue {
	return r.value.PermissionValues
}

func (r *role) SetPermissionValues(pvs []PermissionValue) {
	r.value.PermissionValues = pvs
}

type RoleValue struct {
	// RoleName 角色名
	RoleName string
	// RoleCode 角色编码
	RoleCode string
	// PermissionValues 资源许可列表
	PermissionValues []PermissionValue
}
