package cmd

type TenantCreateCmd struct {
	// 批量新增参数
	Tenants []TenantInfo `json:"tenants"`
}

type TenantInfo struct {
	// 租户ID，非必要，不传默认则随机生成
	TenantID string `json:"tenantID"`
	// 租户名
	TenantName string `json:"tenantName"`
	// 授权码重定向路径，非必要
	RedirectUrl string `json:"redirectUrl"`
	// 访问Token有效时限，非必要，默认7200s
	AccessTokenTimeLimit uint64 `json:"AccessTokenTimeLimit"`
}

type TenantCreateResult struct {
	// 批量响应
	SecretKeys []TenantSecretKey `json:"secretKeys"`
}

type TenantSecretKey struct {
	TenantID  string `json:"tenantID"`
	SecretKey string `json:"secretKey"`
}
