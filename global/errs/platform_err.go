package errs

// 平台层级的错误[2000-2999]
const (
	// TokenNoExistErrorCode token不存在异常
	TokenNoExistErrorCode = 2000 + iota
	// TokenMalformedErrorCode token格式化异常
	TokenMalformedErrorCode
	// TokenInvalidErrorCode token无效
	TokenInvalidErrorCode
	// TokenNotValidYetErrorCode token未生效
	TokenNotValidYetErrorCode
	// TokenExpiredErrorCode token已过期
	TokenExpiredErrorCode
	// PermissionDeniedErrorCode 权限不足拒绝
	PermissionDeniedErrorCode
	// CurrentLimitingErrorCode 平台限流
	CurrentLimitingErrorCode
)

var (
	TokenNoExistErr     = NewUnifiedError(PlatLevel, TokenNoExistErrorCode, "token no exist error")
	TokenMalformedErr   = NewUnifiedError(PlatLevel, TokenMalformedErrorCode, "that's not even a token")
	TokenInvalidErr     = NewUnifiedError(PlatLevel, TokenInvalidErrorCode, "couldn't handle this token")
	TokenNotValidYetErr = NewUnifiedError(PlatLevel, TokenNotValidYetErrorCode, "token not active yet")
	TokenExpiredErr     = NewUnifiedError(PlatLevel, TokenExpiredErrorCode, "token is expired")
	PermissionDeniedErr = NewUnifiedError(PlatLevel, PermissionDeniedErrorCode, "permission denied")
	CurrentLimitingErr  = NewUnifiedError(PlatLevel, CurrentLimitingErrorCode, "platform current limiting")
)

// PlatformErrs 全部平台错误
var PlatformErrs = []*UnifiedError{
	TokenNoExistErr,
	TokenMalformedErr,
	TokenInvalidErr,
	TokenNotValidYetErr,
	TokenExpiredErr,
	PermissionDeniedErr,
	CurrentLimitingErr,
}
