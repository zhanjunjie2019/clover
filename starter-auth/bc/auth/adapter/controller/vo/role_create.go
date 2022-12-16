package vo

type RoleCreateReqVO struct {
	Roles []RoleInfoVO `json:"roles"`
}

type RoleInfoVO struct {
	// 角色名
	RoleName string `json:"roleName" validate:"required"`
	// 角色编码
	RoleCode string `json:"roleCode" validate:"required"`
}

type RoleCreateRspVO struct {
	// 角色ID
	RoleIDs []uint64 `json:"roleIDs"`
}
