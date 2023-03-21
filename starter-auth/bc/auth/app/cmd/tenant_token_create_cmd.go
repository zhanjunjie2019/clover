package cmd

type TenantTokenCreateCmd struct {
	// 租户ID
	TenantID string `json:"tenant_id"`
	// 租户密钥
	SecretKey string `json:"secret_key"`
	// 访问Token过期时间戳，非必要，不传则按当前时间+戳追加租户设置有效时限
	AccessTokenExpirationTime int64 `json:"access_token_expiration_time"`
}

type TenantTokenCreateResult struct {
	// 访问Token
	AccessToken string `json:"access_token"`
	// 访问Token过期时间戳
	AccessTokenExpirationTime int64 `json:"access_token_expiration_time"`
}
