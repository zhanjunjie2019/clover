package vo

type PermissionReqVO struct {
	Permissions []PermissionInfoVO `json:"permissions" validate:"required,gt=0,dive"`
}

type PermissionInfoVO struct {
	// 许可名称
	PermissionName string `json:"permissionName" validate:"required"`
	// 资源编码
	AuthCode string `json:"authCode" validate:"required"`
}

type PermissionRspVO struct {
	// 许可ID
	PermissionIDs []uint64 `json:"permissionIDs"`
}
