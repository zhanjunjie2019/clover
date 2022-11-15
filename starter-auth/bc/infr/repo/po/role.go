package po

import "github.com/zhanjunjie2019/clover/global/defs"

type Role struct {
	defs.ModelPO
	// TenantID 租户ID
	TenantID string `json:"tenantID" gorm:"column:tenant_id;comment:租户ID;size:64;"`
	// RoleName 角色名
	RoleName string `json:"roleName" gorm:"column:role_name;comment:角色名;size:64;"`
	// RoleCode 角色编码
	RoleCode string `json:"roleCode" gorm:"column:role_code;comment:角色编码;size:64;"`
}

func (r Role) TableName() string {
	return "roles_" + r.TenantID
}
