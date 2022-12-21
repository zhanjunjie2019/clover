package cmd

type HelloWorldCmd struct {
	// 姓
	FirstName string `json:"firstName"`
	// 名
	LastName string `json:"lastName"`
}

type HelloWorldResult struct {
	// 响应
	Greetings string
}
