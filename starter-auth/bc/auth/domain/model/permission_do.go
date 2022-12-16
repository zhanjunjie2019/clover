package model

import "github.com/zhanjunjie2019/clover/global/defs"

func NewPermission(id defs.ID, value PermissionValue) Permission {
	return &permission{
		id:    id,
		value: value,
	}
}

type Permission interface {
	ID() defs.ID
	// FullValue 核心数据
	FullValue() PermissionValue
}

type permission struct {
	id    defs.ID
	value PermissionValue
}

func (p permission) ID() defs.ID {
	return p.id
}

func (p permission) FullValue() PermissionValue {
	return p.value
}

type PermissionValue struct {
	// PermissionName 权限名称
	PermissionName string
	// AuthCode 资源编码
	AuthCode string
}
