package po

import "github.com/zhanjunjie2019/clover/global/defs"

type Permission struct {
	defs.ModelPO
	// PermissionName 权限名称
	PermissionName string `gorm:"column:permission_name;comment:权限名称;size:64"`
	// AuthCode 资源编码
	AuthCode string `gorm:"column:auth_code;comment:资源编码;size:64"`
}
