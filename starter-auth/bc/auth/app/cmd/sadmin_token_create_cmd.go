package cmd

type SadminTokenCreateCmd struct {
	// 租户密钥
	SecretKey string `json:"secretKey"`
}

type SadminTokenCreateResult struct {
	// 访问Token
	AccessToken string `json:"accessToken"`
	// 访问Token过期时间戳
	AccessTokenExpirationTime int64 `json:"accessTokenExpirationTime"`
}
