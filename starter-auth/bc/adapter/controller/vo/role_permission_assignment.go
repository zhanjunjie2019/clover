package vo

type RolePermissionAssignmentReqVO struct {
	// 角色编码
	RoleCode string `json:"roleCode" validate:"required"`
	// 资源编码
	AuthCodes []string `json:"authCodes" validate:"required,gt=0,dive,required"`
}

type RolePermissionAssignmentRspVO struct {
	// 角色ID
	RoleId uint64 `json:"roleId"`
}
