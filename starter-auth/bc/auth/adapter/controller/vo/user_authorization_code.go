package vo

type UserAuthorizationCodeReqVO struct {
	// 账户名
	UserName string `json:"user_name" validate:"required"`
	// 密码
	Password string `json:"password" validate:"required"`
}

type UserAuthorizationCodeRspVO struct {
	// 授权码
	AuthorizationCode string `json:"authorization_code"`
	// 重定向路径
	RedirectUrl string `json:"redirect_url"`
}
