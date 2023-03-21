package vo

type UserTokenByAuthcodeReqVO struct {
	// 授权码
	AuthorizationCode string `json:"authorization_code" validate:"required"`
}

type UserTokenByAuthcodeRspVO struct {
	// 访问Token
	AccessToken string `json:"access_token"`
	// 访问Token过期时间戳
	AccessTokenExpirationTime int64 `json:"access_token_expiration_time"`
}
