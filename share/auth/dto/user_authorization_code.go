package dto

import "github.com/zhanjunjie2019/clover/global/defs"

type UserAuthorizationCode struct {
	// TenantID 租户ID
	TenantID string
	// UserID 用户ID
	UserID defs.ID
	// UserName 用户名
	UserName string
}
