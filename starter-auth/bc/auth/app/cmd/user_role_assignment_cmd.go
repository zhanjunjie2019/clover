package cmd

import "github.com/zhanjunjie2019/clover/global/defs"

type UserRoleAssignmentCmd struct {
	// 用户ID，与用户名二选一
	UserID defs.ID `json:"user_id"`
	// 用户名，与用户ID二选一
	UserName string `json:"user_name"`
	// 角色编码
	RoleCodes []string `json:"role_codes"`
}

type UserRoleAssignmentResult struct {
	// 用户ID
	UserID defs.ID `json:"user_id"`
}
