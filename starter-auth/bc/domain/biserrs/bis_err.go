package biserrs

import "github.com/zhanjunjie2019/clover/global/errs"

// 租户的错误[5000-5999]
const (
	// TenantAlreadyExistErrorCode 租户已存在
	TenantAlreadyExistErrorCode = 5000 + iota
	// TenantDoesNotExistErrorCode 租户不在
	TenantDoesNotExistErrorCode
	// PermissionAlreadyExistErrorCode 权限已存在
	PermissionAlreadyExistErrorCode
	// TenantValidationFailedErrorCode 租户校验失败
	TenantValidationFailedErrorCode
	// UserAlreadyExistsErrorCode 用户已存在
	UserAlreadyExistsErrorCode
	// LoginVerifyFailedErrorCode 登录验证失败
	LoginVerifyFailedErrorCode
	// RoleAlreadyExistsErrorCode 角色已存在
	RoleAlreadyExistsErrorCode
)

var (
	// TenantAlreadyExistErr 租户已存在
	TenantAlreadyExistErr = errs.NewUnifiedError(errs.BisLevel, TenantAlreadyExistErrorCode, "tenant already exist")
	// TenantDoesNotExistErr 租户不在
	TenantDoesNotExistErr = errs.NewUnifiedError(errs.BisLevel, TenantDoesNotExistErrorCode, "tenant does not exist")
	// PermissionAlreadyExistErr 权限已存在
	PermissionAlreadyExistErr = errs.NewUnifiedError(errs.BisLevel, PermissionAlreadyExistErrorCode, "permission already exist")
	// TenantValidationFailedErr 租户校验失败
	TenantValidationFailedErr = errs.NewUnifiedError(errs.BisLevel, TenantValidationFailedErrorCode, "tenant validation failed")
	// UserAlreadyExistsErr 用户已存在
	UserAlreadyExistsErr = errs.NewUnifiedError(errs.BisLevel, UserAlreadyExistsErrorCode, "user already exist")
	// LoginVerifyFailedErr 登录验证失败
	LoginVerifyFailedErr = errs.NewUnifiedError(errs.BisLevel, LoginVerifyFailedErrorCode, "login verification failed")
	// RoleAlreadyExistsErr 角色已存在
	RoleAlreadyExistsErr = errs.NewUnifiedError(errs.BisLevel, RoleAlreadyExistsErrorCode, "role already exist")
)
