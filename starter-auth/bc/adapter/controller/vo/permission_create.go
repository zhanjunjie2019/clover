package vo

type PermissionReqVO struct {
	// 许可名称
	PermissionName string
	// 资源编码
	AuthCode string
}

type PermissionRspVO struct {
	// 许可ID
	PermissionId uint64 `json:"permissionId"`
}
