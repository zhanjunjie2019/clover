package vo

type UserRoleAssignmentReqVO struct {
	// 用户ID，与用户名二选一
	UserID uint64 `json:"userID" validate:"required_without=UserName"`
	// 用户名，与用户ID二选一
	UserName string `json:"userName" validate:"required_without=UserID"`
	// 角色编码
	RoleCodes []string `json:"roleCodes" validate:"required,gt=0,dive,required"`
}

type UserRoleAssignmentRspVO struct {
	// 用户ID
	UserID uint64 `json:"userID"`
}
