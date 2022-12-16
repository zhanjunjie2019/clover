package vo

type SadminTokenCreateReqVO struct {
	// 租户密钥
	SecretKey string `json:"secretKey" validate:"required"`
}

type SadminTokenCreateRspVO struct {
	// 访问Token
	AccessToken string `json:"accessToken"`
	// 访问Token过期时间戳
	AccessTokenExpirationTime int64 `json:"accessTokenExpirationTime"`
}
