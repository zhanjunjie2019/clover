package vo

type SadminTokenCreateReqVO struct {
	// 租户密钥
	SecretKey string `json:"secret_key" validate:"required"`
}

type SadminTokenCreateRspVO struct {
	// 访问Token
	AccessToken string `json:"access_token"`
	// 访问Token过期时间戳
	AccessTokenExpirationTime int64 `json:"access_token_expiration_time"`
}
