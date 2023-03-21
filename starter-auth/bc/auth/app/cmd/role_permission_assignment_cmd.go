package cmd

import "github.com/zhanjunjie2019/clover/global/defs"

type RolePermissionAssignmentCmd struct {
	// 角色ID，与角色编码二选一
	RoleID defs.ID `json:"role_id"`
	// 角色编码，与角色ID二选一
	RoleCode string `json:"role_code"`
	// 资源编码
	AuthCodes []string `json:"auth_codes"`
}

type RolePermissionAssignmentResult struct {
	// 角色ID
	RoleID defs.ID `json:"role_id"`
}
