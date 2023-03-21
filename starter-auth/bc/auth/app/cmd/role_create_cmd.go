package cmd

import "github.com/zhanjunjie2019/clover/global/defs"

type RoleCreateCmd struct {
	Roles []RoleInfo `json:"roles"`
}

type RoleInfo struct {
	// 角色名
	RoleName string `json:"role_name"`
	// 角色编码
	RoleCode string `json:"role_code"`
}

type RoleCreateResult struct {
	// 角色ID
	RoleIDs []defs.ID `json:"role_ids"`
}
