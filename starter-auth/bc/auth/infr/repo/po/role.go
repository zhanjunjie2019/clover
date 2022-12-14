package po

import "github.com/zhanjunjie2019/clover/global/defs"

type Role struct {
	defs.ModelPO
	// RoleName 角色名
	RoleName string `gorm:"column:role_name;comment:角色名;size:64"`
	// RoleCode 角色编码
	RoleCode string `gorm:"column:role_code;comment:角色编码;size:64"`
}
