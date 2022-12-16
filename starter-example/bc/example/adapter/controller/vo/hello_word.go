package vo

type HelloWordReqVO struct {
	// 姓，必填
	FirstName string `json:"firstName" form:"firstName" validate:"required"`
	// 名，必填
	LastName string `json:"lastName" form:"lastName" validate:"required"`
}

type HelloWordRspVO struct {
	// 问候语
	Greetings string `json:"greetings"`
}
