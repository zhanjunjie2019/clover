package cmd

import "github.com/zhanjunjie2019/clover/global/defs"

type RolePermissionAssignmentCmd struct {
	// 角色ID，与角色编码二选一
	RoleID defs.ID `json:"roleID"`
	// 角色编码，与角色ID二选一
	RoleCode string `json:"roleCode"`
	// 资源编码
	AuthCodes []string `json:"authCodes"`
}

type RolePermissionAssignmentResult struct {
	// 角色ID
	RoleID defs.ID `json:"roleID"`
}
