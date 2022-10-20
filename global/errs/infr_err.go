package errs

// 底层的错误[1000-1999]
const (
	// UnknownErrorCode 未知错误
	UnknownErrorCode = 1000 + iota
	// FieldKindTypeErrorCode 字段类型不正确
	FieldKindTypeErrorCode
	// DBLinkTypeErrorCode 数据库类型不匹配
	DBLinkTypeErrorCode
	// DBConnectionErrorCode 数据库连接失败
	DBConnectionErrorCode
)

var (
	// UnknownErr 未知错误
	UnknownErr = NewUnifiedError(InfrLevel, UnknownErrorCode, "unknown error")
	// FieldKindTypeErr 字段类型不正确
	FieldKindTypeErr = NewUnifiedError(InfrLevel, FieldKindTypeErrorCode, "file kind type error")
	// DBLinkTypeErr 数据库类型不匹配
	DBLinkTypeErr = NewUnifiedError(InfrLevel, DBLinkTypeErrorCode, "db link type error")
	// DBConnectionErr 数据库连接失败
	DBConnectionErr = NewUnifiedError(InfrLevel, DBConnectionErrorCode, "db connection error")
)

func NewUnknownErr(err error) *UnifiedError {
	return NewUnifiedError(InfrLevel, UnknownErrorCode, "unknown error:"+err.Error())
}

// InfrErrs 全部底层错误
var InfrErrs = []*UnifiedError{
	UnknownErr,
	FieldKindTypeErr,
	DBLinkTypeErr,
	DBConnectionErr,
}
