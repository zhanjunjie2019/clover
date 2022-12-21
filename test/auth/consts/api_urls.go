package consts

import "os"

var DomainHost string

func init() {
	domainHost := os.Getenv("DOMAIN_HOST")
	if len(domainHost) > 0 {
		DomainHost = domainHost
	} else {
		DomainHost = "http://localhost:8810"
	}
}

const (
	SadminTokenCreateApiUrl        = "/auth/sadmin-token-create"
	PermissionCreateApiUrl         = "/auth/permission-create"
	TenantCreateApiUrl             = "/auth/tenant-create"
	TenantTokenCreateApiUrl        = "/auth/tenant-token-create"
	RoleCreateApiUrl               = "/auth/role-create"
	RolePermissionAssignmentApiUrl = "/auth/role-permission-assignment"
	UserCreateApiUrl               = "/auth/user-create"
	UserRoleAssignmentApiUrl       = "/auth/user-role-assignment"
	UserAuthorizationCodeApiUrl    = "/auth/user-authorization-code"
)
