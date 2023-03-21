package defs

import (
	"github.com/golang-jwt/jwt/v4"
)

type JwtClaims struct {
	jwt.RegisteredClaims
	// TenantID 租户ID
	TenantID string `json:"tenant_id,omitempty"`
	// UserID 用户ID
	UserID uint64 `json:"user_id,omitempty"`
	// Username 用户名
	Username string `json:"username,omitempty"`
	// Auths 资源许可
	Auths []string `json:"auths,omitempty"`
}
