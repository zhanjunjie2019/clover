package vo

type UserCreateReqVO struct {
	// UserName 用户名
	UserName string `json:"userName" validate:"required"`
	// Password 密码
	Password string `json:"password" validate:"required"`
}

type UserCreateRspVO struct {
	// UserId 用户ID
	UserId uint64 `json:"userId"`
}
