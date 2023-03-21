package vo

type RoleCreateReqVO struct {
	Roles []RoleInfoVO `json:"roles" validate:"required,gt=0,dive"`
}

type RoleInfoVO struct {
	// 角色名
	RoleName string `json:"role_name" validate:"required"`
	// 角色编码
	RoleCode string `json:"role_code" validate:"required"`
}

type RoleCreateRspVO struct {
	// 角色ID
	RoleIDs []uint64 `json:"role_ids"`
}
