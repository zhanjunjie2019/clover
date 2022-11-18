package vo

type TenantTokenCreateReqVO struct {
	// 租户ID
	TenantID string `json:"tenantID" validate:"required"`
	// 租户密钥
	SecretKey string `json:"secretKey" validate:"required"`
	// 访问Token过期时间戳，非必要，不传则按当前时间+戳追加租户设置有效时限
	AccessTokenExpirationTime int64 `json:"accessTokenExpirationTime"`
	// 刷新Token过期时间戳，非必要，不传则按当前时间+戳追加租户设置有效时限
	RefreshTokenExpirationTime int64 `json:"refreshTokenExpirationTime"`
}

type TenantTokenCreateRspVO struct {
	// 访问Token
	AccessToken string `json:"accessToken"`
	// 访问Token过期时间戳
	AccessTokenExpirationTime int64 `json:"accessTokenExpirationTime"`
	// 刷新Token
	RefreshToken string `json:"refreshToken"`
	// 刷新Token过期时间戳
	RefreshTokenExpirationTime int64 `json:"refreshTokenExpirationTime"`
}
