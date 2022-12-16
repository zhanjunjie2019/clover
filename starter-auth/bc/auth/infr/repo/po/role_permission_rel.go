package po

import "github.com/zhanjunjie2019/clover/global/defs"

type RolePermissionRel struct {
	defs.ModelPO
	// RoleId 角色ID
	RoleId defs.ID `gorm:"column:role_id;comment:角色ID"`
	// RoleCode 角色编码
	RoleCode string `gorm:"column:role_code;comment:角色编码;size:64"`
	// PermissionName 权限名称
	PermissionName string `gorm:"column:permission_name;comment:权限名称;size:64"`
	// AuthCode 资源编码
	AuthCode string `gorm:"column:auth_code;comment:资源编码;size:64"`
}
