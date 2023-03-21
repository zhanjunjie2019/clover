package vo

type PermissionReqVO struct {
	Permissions []PermissionInfoVO `json:"permissions" validate:"required,gt=0,dive"`
}

type PermissionInfoVO struct {
	// 许可名称
	PermissionName string `json:"permission_name" validate:"required"`
	// 资源编码
	AuthCode string `json:"auth_code" validate:"required"`
}

type PermissionRspVO struct {
	// 许可ID
	PermissionIDs []uint64 `json:"permission_ids"`
}
