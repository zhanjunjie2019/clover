package dto

type TenantInitEventDTO struct {
	// TenantID 租户ID
	TenantID string `json:"tenantID"`
	// TenantName 租户名
	TenantName string `json:"tenantName"`
}
