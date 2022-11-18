package biserrs

import "github.com/zhanjunjie2019/clover/global/errs"

// 租户的错误[5000-5999]
const (
	// TenantAlreadyExistErrorCode 租户已存在
	TenantAlreadyExistErrorCode = 5000 + iota
	// LoginVerifyFailedErrorCode 登录验证失败
	LoginVerifyFailedErrorCode
)

var (
	// TenantAlreadyExistErr 租户已存在
	TenantAlreadyExistErr = errs.NewUnifiedError(errs.BisLevel, TenantAlreadyExistErrorCode, "tenant already exist")
	// LoginVerifyFailedErr 登录验证失败
	LoginVerifyFailedErr = errs.NewUnifiedError(errs.BisLevel, LoginVerifyFailedErrorCode, "login verification failed")
)
