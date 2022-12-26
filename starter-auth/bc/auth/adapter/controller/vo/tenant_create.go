package vo

type TenantCreateReqVO struct {
	Tenants []TenantInfoVO `json:"tenants" validate:"required,gt=0,dive"`
}

type TenantInfoVO struct {
	// 租户ID，非必要，不传默认则随机生成，英文字母，长度小于{20}位
	TenantID string `json:"tenantID" validate:"lte=20"`
	// 租户名
	TenantName string `json:"tenantName" validate:"required"`
	// 授权码重定向路径，非必要
	RedirectUrl string `json:"redirectUrl"`
	// 访问Token有效时限，非必要，默认7200s
	AccessTokenTimeLimit uint64 `json:"AccessTokenTimeLimit"`
}

type TenantCreateRspVO struct {
	SecretKeys []TenantSecretKeyVO `json:"secretKeys"`
}

type TenantSecretKeyVO struct {
	// 租户ID
	TenantID string `json:"tenantID"`
	// 租户密钥
	SecretKey string `json:"secretKey"`
}
