package cmd

type HelloWorldCmd struct {
	// 姓
	FirstName string `json:"first_name"`
	// 名
	LastName string `json:"last_name"`
}

type HelloWorldResult struct {
	// 响应
	Greetings string `json:"greetings"`
}
