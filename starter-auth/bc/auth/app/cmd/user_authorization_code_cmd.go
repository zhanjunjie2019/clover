package cmd

type UserAuthorizationCodeCmd struct {
	// 账户名
	UserName string `json:"user_name"`
	// 密码
	Password string `json:"password"`
}

type UserAuthorizationCodeResult struct {
	// 授权码
	AuthorizationCode string `json:"authorization_code"`
	// 重定向路径
	RedirectUrl string `json:"redirect_url"`
}
