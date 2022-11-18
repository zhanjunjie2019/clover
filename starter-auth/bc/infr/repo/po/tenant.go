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
}
