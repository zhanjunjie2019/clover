package vo

type TenantCreateReqVO struct {
	// 租户ID，不传默认则随机生成
	TenantID string `json:"tenantID"`
	// 租户名
	TenantName string `json:"tenantName" validate:"required"`
}

type TenantCreateRspVO struct {
	// 租户ID
	TenantID string `json:"tenantID"`
	// 租户密钥
	SecretKey string `json:"secretKey"`
}
