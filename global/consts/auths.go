package consts

const (
	// SAdminAuth 平台超管，处理跨租户业务
	SAdminAuth = "SADMIN"
	// AdminAuth 租户管理员，处理单一租户业务，必须在租户Token中使用
	AdminAuth = "ADMIN"
)
