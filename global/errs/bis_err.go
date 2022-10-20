package errs

// 业务的错误[4000-4999]
const (
	// ReqParamsErrorCode 参数错误
	ReqParamsErrorCode = 4000 + iota
)

var (
	ReqParamsErr = NewUnifiedError(BisLevel, ReqParamsErrorCode, "request params error")
)
