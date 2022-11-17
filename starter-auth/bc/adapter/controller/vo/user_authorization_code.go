package vo

type UserAuthorizationCodeReqVO struct {
	// 账户名
	UserName string `json:"userName" validate:"required"`
	// 密码
	Password string `json:"password" validate:"required"`
	// 重定向路径，可选，若无按租户设置
	RedirectUrl string `json:"redirectUrl"`
}

type UserAuthorizationCodeRspVO struct {
	// 授权码
	AuthorizationCode string `json:"authorizationCode"`
	// 重定向路径
	RedirectUrl string `json:"redirectUrl"`
}
