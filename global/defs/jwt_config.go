package defs

import (
	"github.com/golang-jwt/jwt/v4"
)

type JwtClaims struct {
	jwt.RegisteredClaims
	// TenantID 租户ID
	TenantID string `json:"tenantID,omitempty"`
	// UserID 用户ID
	UserID uint64 `json:"userID,omitempty"`
	// Username 用户名
	Username string `json:"username,omitempty"`
	// Auths 资源权限
	Auths []string `json:"auths,omitempty"`
}
