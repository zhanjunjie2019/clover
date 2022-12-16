package po

import "github.com/zhanjunjie2019/clover/global/defs"

type UserRoleRel struct {
	defs.ModelPO
	// UserID 用户ID
	UserID defs.ID `gorm:"column:user_id;comment:用户ID"`
	// UserName 用户名
	UserName string `gorm:"column:user_name;comment:用户名;size:64"`
	// RoleName 角色名
	RoleName string `gorm:"column:role_name;comment:角色名;size:64"`
	// RoleCode 角色编码
	RoleCode string `gorm:"column:role_code;comment:角色编码;size:64"`
}
