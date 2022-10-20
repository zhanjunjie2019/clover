package biserrs

import "github.com/zhanjunjie2019/clover/global/errs"

// 租户的错误[5000-5999]
const (
	// TenantAlreadyExistErrorCode 租户已存在
	TenantAlreadyExistErrorCode = 5000 + iota
)

var (
	// TenantAlreadyExistErr 租户已存在
	TenantAlreadyExistErr = errs.NewUnifiedError(errs.BisLevel, TenantAlreadyExistErrorCode, "tenant already exist")
)
