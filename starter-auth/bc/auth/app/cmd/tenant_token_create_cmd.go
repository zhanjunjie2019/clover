package cmd

type TenantTokenCreateCmd struct {
	// 租户ID
	TenantID string `json:"tenantID"`
	// 租户密钥
	SecretKey string `json:"secretKey"`
	// 访问Token过期时间戳，非必要，不传则按当前时间+戳追加租户设置有效时限
	AccessTokenExpirationTime int64 `json:"accessTokenExpirationTime"`
}

type TenantTokenCreateResult struct {
	// 访问Token
	AccessToken string `json:"accessToken"`
	// 访问Token过期时间戳
	AccessTokenExpirationTime int64 `json:"accessTokenExpirationTime"`
}
