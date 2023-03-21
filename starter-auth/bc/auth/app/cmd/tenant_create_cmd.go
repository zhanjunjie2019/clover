package cmd

type TenantCreateCmd struct {
	// 批量新增参数
	Tenants []TenantInfo `json:"tenants"`
}

type TenantInfo struct {
	// 租户ID，非必要，不传默认则随机生成
	TenantID string `json:"tenant_id"`
	// 租户名
	TenantName string `json:"tenant_name"`
	// 授权码重定向路径，非必要
	RedirectUrl string `json:"redirect_url"`
	// 访问Token有效时限，非必要，默认7200s
	AccessTokenTimeLimit uint64 `json:"access_token_time_limit"`
}

type TenantCreateResult struct {
	// 批量响应
	SecretKeys []TenantSecretKey `json:"secret_keys"`
}

type TenantSecretKey struct {
	TenantID  string `json:"tenant_id"`
	SecretKey string `json:"secret_key"`
}
