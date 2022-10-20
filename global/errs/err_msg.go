package errs

import "fmt"

type ErrLevel int

const (
	InfrLevel ErrLevel = 1 + iota
	PlatLevel
	BisLevel
)

const (
	defaultBisErrorMsg = "system error"
)

func ToUnifiedError(err error) *UnifiedError {
	if unifiedErr, ok := err.(*UnifiedError); ok {
		return unifiedErr
	} else {
		return NewUnknownErr(err)
	}
}

func NewUnifiedError(level ErrLevel, code int, msg string) *UnifiedError {
	return &UnifiedError{level: level, code: code, msg: msg}
}

type UnifiedError struct {
	level ErrLevel
	code  int
	msg   string
}

func (u UnifiedError) Code() int {
	return u.code
}

func (u UnifiedError) Error() string {
	return fmt.Sprintf("[%d]%s", u.code, u.msg)
}

func (u UnifiedError) ShowError() string {
	if u.level == InfrLevel {
		return fmt.Sprintf("[%d]%s", u.code, defaultBisErrorMsg)
	}
	return fmt.Sprintf("[%d]%s", u.code, u.msg)
}
