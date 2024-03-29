package cmd

import "github.com/zhanjunjie2019/clover/global/defs"

type PermissionCreateCmd struct {
	Permissions []PermissionInfo `json:"permissions"`
}

// PermissionInfo 资源许可
type PermissionInfo struct {
	// 许可名称
	PermissionName string `json:"permission_name"`
	// 资源编码
	AuthCode string `json:"auth_code"`
}

type PermissionCreateResult struct {
	PermissionIDs []defs.ID `json:"permission_ids"`
}
