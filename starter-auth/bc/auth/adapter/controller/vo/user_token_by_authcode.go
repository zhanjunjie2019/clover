package vo

type UserTokenByAuthcodeReqVO struct {
	// 授权码
	AuthorizationCode string `json:"authorizationCode" validate:"required"`
}

type UserTokenByAuthcodeRspVO struct {
	// 访问Token
	AccessToken string `json:"accessToken"`
	// 访问Token过期时间戳
	AccessTokenExpirationTime int64 `json:"accessTokenExpirationTime"`
}
