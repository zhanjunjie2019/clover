package vo

type PermissionReqVO struct {
	// 许可名称
	PermissionName string `json:"permissionName" validate:"required"`
	// 资源编码
	AuthCode string `json:"authCode" validate:"required"`
}

type PermissionRspVO struct {
	// 许可ID
	PermissionID uint64 `json:"permissionID"`
}
