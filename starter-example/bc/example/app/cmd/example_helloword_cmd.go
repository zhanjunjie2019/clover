package cmd

type HelloWordCmd struct {
	// 姓
	FirstName string `json:"firstName"`
	// 名
	LastName string `json:"lastName"`
}

type HelloWordResult struct {
	// 响应
	Greetings string
}
