package vo

type RolePermissionAssignmentReqVO struct {
	// 角色ID，与角色编码二选一
	RoleID uint64 `json:"roleID" validate:"required_without=RoleCode"`
	// 角色编码，与角色ID二选一
	RoleCode string `json:"roleCode" validate:"required_without=RoleID"`
	// 资源编码
	AuthCodes []string `json:"authCodes" validate:"required,gt=0,dive,required"`
}

type RolePermissionAssignmentRspVO struct {
	// 角色ID
	RoleID uint64 `json:"roleID"`
}
