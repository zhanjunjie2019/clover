package po

import "github.com/zhanjunjie2019/clover/global/defs"

type Tenant struct {
	defs.ModelPO
	// TenantID 租户ID
	TenantID string `json:"tenantID" gorm:"column:tenant_id;comment:租户ID;size:64;"`
	// TenantName 租户名
	TenantName string `json:"tenantName" gorm:"column:tenant_name;comment:租户名;size:64;"`
	// SecretKey 租户密钥
	SecretKey string `json:"secretKey" gorm:"column:secret_key;comment:租户密钥;size:64;"`
	// RedirectUrl 授权码定向路径
	RedirectUrl string `json:"redirectUrl" gorm:"column:redirect_url;comment:授权码定向路径;size:512;"`
	// AccessTokenTimeLimit 访问Token有效时限
	AccessTokenTimeLimit uint64 `json:"accessTokenTimeLimit" gorm:"column:access_token_time_timit;comment:访问Token有效时限;"`
	// RefreshTokenTimeLimit 刷新Token有效时限
	RefreshTokenTimeLimit uint64 `json:"refreshTokenTimeLimit" gorm:"column:refresh_token_time_timit;comment:刷新Token有效时限;"`
}
