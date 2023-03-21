package vo

type HelloWorldReqVO struct {
	// 姓，必填
	FirstName string `json:"first_name" form:"firstName" validate:"required"`
	// 名，必填
	LastName string `json:"last_name" form:"lastName" validate:"required"`
}

type HelloWorldRspVO struct {
	// 问候语
	Greetings string `json:"greetings"`
}
