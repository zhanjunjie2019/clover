package cmd

type HelloWorldCmd struct {
	// ε§
	FirstName string `json:"firstName"`
	// ε
	LastName string `json:"lastName"`
}

type HelloWorldResult struct {
	// εεΊ
	Greetings string
}
