package vo

type UserCreateReqVO struct {
	Users []UserInfoVO `json:"users" validate:"required,gt=0,dive"`
}

type UserInfoVO struct {
	// 用户名
	UserName string `json:"userName" validate:"required"`
	// 密码
	Password string `json:"password" validate:"required"`
}

type UserCreateRspVO struct {
	// 用户ID
	UserIDs []uint64 `json:"userIDs"`
}
