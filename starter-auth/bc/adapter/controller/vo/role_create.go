package vo

type RoleCreateReqVO struct {
	// 角色名
	RoleName string `json:"roleName" validate:"required"`
	// 角色编码
	RoleCode string `json:"roleCode" validate:"required"`
}

type RoleCreateRspVO struct {
	// 角色ID
	RoleID uint64 `json:"roleID"`
}
