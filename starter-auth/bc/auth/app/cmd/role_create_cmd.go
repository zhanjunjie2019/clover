package cmd

import "github.com/zhanjunjie2019/clover/global/defs"

type RoleCreateCmd struct {
	Roles []RoleInfo `json:"roles"`
}

type RoleInfo struct {
	// 角色名
	RoleName string `json:"roleName"`
	// 角色编码
	RoleCode string `json:"roleCode"`
}

type RoleCreateResult struct {
	// 许可ID
	PermissionIDs []defs.ID `json:"roleIDs"`
}
