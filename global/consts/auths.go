package consts

const (
	// SadminAuth 平台超管，处理跨租户业务
	SadminAuth = "SADMIN"
	// AdminAuth 租户管理员，处理单一租户业务，必须配合租户Token使用
	AdminAuth = "ADMIN"
)
