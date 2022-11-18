package vo

type TenantCreateReqVO struct {
	// 租户ID，非必要，不传默认则随机生成
	TenantID string `json:"tenantID"`
	// 租户名
	TenantName string `json:"tenantName" validate:"required"`
	// 授权码重定向路径，非必要
	RedirectUrl string `json:"redirectUrl"`
	// 访问Token有效时限，非必要，默认7200s
	AccessTokenTimeLimit uint64 `json:"AccessTokenTimeLimit"`
	// 刷新Token有效时限，非必要，默认86400s
	RefreshTokenTimeLimit uint64 `json:"refreshTokenTimeLimit"`
}

type TenantCreateRspVO struct {
	// 租户ID
	TenantID string `json:"tenantID"`
	// 租户密钥
	SecretKey string `json:"secretKey"`
}
