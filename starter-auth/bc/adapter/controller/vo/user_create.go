package vo

type UserCreateReqVO struct {
	// 用户名
	UserName string `json:"userName" validate:"required"`
	// 密码
	Password string `json:"password" validate:"required"`
}

type UserCreateRspVO struct {
	// 用户ID
	UserId uint64 `json:"userId"`
}
