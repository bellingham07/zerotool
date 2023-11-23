package errorx

import (
	"fmt"
	"github.com/bellingham07/gozero-tools/codex"
	"net/http"
)

type CodeError struct {
	Code     int      `json:"code"`
	Msg      string   `json:"msg"`
	BizType  string   `json:"-"` // 业务类型
	Metadata Metadata `json:"-"` // 错误元数据
	Err      error    `json:"-"`
}
type Metadata map[string]interface{}

func New(t string, code int, message string) *CodeError {
	return &CodeError{
		BizType: t,
		Code:    code,
		Msg:     message,
	}
}

func WithCode(t string, code codex.ResCode) *CodeError {
	return &CodeError{
		BizType: t,
		Code:    int(code),
		Msg:     code.Msg(),
	}
}

func Internal(err error, format string, args ...interface{}) *CodeError {
	message := format
	if len(args) > 0 {
		message = fmt.Sprintf(format, args...)
	}
	return New(http.StatusText(http.StatusInternalServerError),
		http.StatusInternalServerError, message).WithError(err)
}

// 实现了error的Error()方法，CodeError就是一个error
func (e *CodeError) Error() string {
	if e.Err != nil {
		return e.Msg + ": " + e.Err.Error()
	}
	return e.Msg
}
func (e *CodeError) WithError(err error) *CodeError {
	e.Err = err
	return e
}
func (e *CodeError) WithMetadata(metadata Metadata) *CodeError {
	e.Metadata = metadata
	return e
}
func From(err error) *CodeError {
	if err == nil {
		return nil
	}
	if errx, ok := err.(*CodeError); ok {
		return errx
	}
	return Internal(err, codex.CodeInternalErr.Msg())
}
