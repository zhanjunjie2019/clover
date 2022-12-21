package cmd

type UserTokenByAuthcodeCmd struct {
	// 授权码
	AuthorizationCode string `json:"authorizationCode"`
}

type UserTokenByAuthcodeResult struct {
	// 访问Token
	AccessToken string `json:"accessToken"`
	// 访问Token过期时间戳
	AccessTokenExpirationTime int64 `json:"accessTokenExpirationTime"`
}
