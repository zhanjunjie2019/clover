package vo

type TenantCreateReqVO struct {
	Tenants []TenantInfoVO `json:"tenants" validate:"required,gt=0,dive"`
}

type TenantInfoVO struct {
	// 租户ID，非必要，不传默认则随机生成，英文字母，长度小于{20}位
	TenantID string `json:"tenant_id" validate:"lte=20"`
	// 租户名
	TenantName string `json:"tenant_name" validate:"required"`
	// 授权码重定向路径，非必要
	RedirectUrl string `json:"redirect_url"`
	// 访问Token有效时限，非必要，默认7200s
	AccessTokenTimeLimit uint64 `json:"access_token_time_limit"`
}

type TenantCreateRspVO struct {
	SecretKeys []TenantSecretKeyVO `json:"secret_keys"`
}

type TenantSecretKeyVO struct {
	// 租户ID
	TenantID string `json:"tenant_id"`
	// 租户密钥
	SecretKey string `json:"secret_key"`
}
