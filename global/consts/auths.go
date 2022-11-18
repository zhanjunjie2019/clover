package consts

const (
	// SAdminAuth 平台超管，处理跨租户业务
	SAdminAuth = "SADMIN"
	// AdminAuth 租户管理员，处理单一租户业务，必须在租户Token中使用
	AdminAuth = "ADMIN"
	// RefreshAdminAuth 租户管理员,同于刷新获得新的Token
	RefreshAdminAuth = "REFRESH_ADMIN"
)
