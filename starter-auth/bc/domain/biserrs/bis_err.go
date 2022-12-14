package biserrs

import (
	"fmt"
	"github.com/zhanjunjie2019/clover/global/errs"
)

// 租户的错误[5000-5999]
const (
	// TenantAlreadyExistErrorCode 租户已存在
	TenantAlreadyExistErrorCode = 5000 + iota
	// TenantDoesNotExistErrorCode 租户不存在
	TenantDoesNotExistErrorCode
	// PermissionAlreadyExistErrorCode 权限已存在
	PermissionAlreadyExistErrorCode
	// PermissionDoesNotExistErrorCode 权限不存在
	PermissionDoesNotExistErrorCode
	// TenantValidationFailedErrorCode 租户校验失败
	TenantValidationFailedErrorCode
	// UserAlreadyExistsErrorCode 用户已存在
	UserAlreadyExistsErrorCode
	// UserDoesNotExistErrorCode 用户不存在
	UserDoesNotExistErrorCode
	// LoginVerifyFailedErrorCode 登录验证失败
	LoginVerifyFailedErrorCode
	// RoleAlreadyExistsErrorCode 角色已存在
	RoleAlreadyExistsErrorCode
	// RoleDoesNotExistErrorCode 角色不存在
	RoleDoesNotExistErrorCode
)

var (
	// TenantValidationFailedErr 租户校验失败
	TenantValidationFailedErr = errs.NewUnifiedError(errs.BisLevel, TenantValidationFailedErrorCode, "tenant validation failed")
	// LoginVerifyFailedErr 登录验证失败
	LoginVerifyFailedErr = errs.NewUnifiedError(errs.BisLevel, LoginVerifyFailedErrorCode, "login verification failed")
	// UserDoesNotExistErr 用户不存在
	UserDoesNotExistErr = errs.NewUnifiedError(errs.BisLevel, UserDoesNotExistErrorCode, "user does not exist")
	// RoleDoesNotExistErr 角色不存在
	RoleDoesNotExistErr = errs.NewUnifiedError(errs.BisLevel, RoleDoesNotExistErrorCode, "role does not exist")
)

// TenantAlreadyExistErrWithTenantID 租户已存在
func TenantAlreadyExistErrWithTenantID(tenantID string) *errs.UnifiedError {
	return errs.NewUnifiedError(errs.BisLevel, TenantAlreadyExistErrorCode, fmt.Sprintf("tenant(tenantID=%s) already exist", tenantID))
}

// TenantDoesNotExistErrWithTenantID 租户不在
func TenantDoesNotExistErrWithTenantID(tenantID string) *errs.UnifiedError {
	return errs.NewUnifiedError(errs.BisLevel, TenantDoesNotExistErrorCode, fmt.Sprintf("tenant(tenantID=%s) does not exist", tenantID))
}

// PermissionAlreadyExistErrWithAuthCode 权限已存在
func PermissionAlreadyExistErrWithAuthCode(authCode string) *errs.UnifiedError {
	return errs.NewUnifiedError(errs.BisLevel, PermissionAlreadyExistErrorCode, fmt.Sprintf("permission(authCode=%s) already exist", authCode))
}

// PermissionDoesNotExistErrWithAuthCode 权限不存在
func PermissionDoesNotExistErrWithAuthCode(authCode string) *errs.UnifiedError {
	return errs.NewUnifiedError(errs.BisLevel, PermissionDoesNotExistErrorCode, fmt.Sprintf("permission(authCode=%s) does not exist", authCode))
}

// UserAlreadyExistsErrWithUserName 用户已存在
func UserAlreadyExistsErrWithUserName(userName string) *errs.UnifiedError {
	return errs.NewUnifiedError(errs.BisLevel, UserAlreadyExistsErrorCode, fmt.Sprintf("user(userName=%s) already exist", userName))
}

// RoleAlreadyExistsErrWithRoleCode 角色已存在
func RoleAlreadyExistsErrWithRoleCode(roleCode string) *errs.UnifiedError {
	return errs.NewUnifiedError(errs.BisLevel, RoleAlreadyExistsErrorCode, fmt.Sprintf("role(roleCode=%s) already exist", roleCode))
}

// RoleDoesNotExistErrWithRoleCode 角色不存在
func RoleDoesNotExistErrWithRoleCode(roleCode string) *errs.UnifiedError {
	return errs.NewUnifiedError(errs.BisLevel, RoleDoesNotExistErrorCode, fmt.Sprintf("role(roleCode=%s) does not exist", roleCode))
}
