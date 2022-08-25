package error_utils

import (
	"fmt"
)

const (
	BizErrorCodeSuccess = 0
	BizErrorCodeError   = 1
)

type BizError struct {
	Code         int
	Message      string
	OtherMessage string
	ActualErr    error
	Entity       interface{}
}

func NewBizError(code int, msg string, others ...string) BizError {
	otherMsg := ""
	if len(others) > 0 {
		otherMsg = others[0]
	}
	return BizError{code, msg, otherMsg, nil, nil}
}

func BizErrorWrapper(err error) (bizErr BizError) {
	var ok bool
	if bizErr, ok = err.(BizError); !ok {
		bizErr = BizError{
			Code:      BizErrorCodeError,
			ActualErr: err,
		}
	}

	return
}

func (e BizError) WithOtherMessage(format string, arg ...interface{}) BizError {
	e.OtherMessage = fmt.Sprintf("%s %s", e.OtherMessage, fmt.Sprintf(format, arg...))
	return e
}

func (e BizError) WithActualErr(err error) BizError {
	e.ActualErr = err
	return e
}

func (e BizError) WithActualErrAndMessage(err error, format string, arg ...interface{}) BizError {
	e.WithOtherMessage(format, arg...)
	e.ActualErr = err
	return e
}

func (e BizError) WithEntity(entity interface{}) BizError {
	e.Entity = entity
	return e
}

// 实现 Error() 方法, 即实现 error 接口
func (e BizError) Error() (msg string) {
	if e.OtherMessage != "" {
		msg = fmt.Sprintf("%s %s", e.Message, e.OtherMessage)
	} else {
		msg = e.Message
	}
	return
}

func (e BizError) Details() (msg string) {
	msg = e.Message

	if e.ActualErr != nil {
		msg = fmt.Sprintf("%s %s", msg, e.ActualErr.Error())
	}

	if e.OtherMessage != "" {
		msg = fmt.Sprintf("%s %s", msg, e.OtherMessage)
	}

	return
}

var (
	//成功
	BizErrorSuccess = NewBizError(BizErrorCodeSuccess, "Success.")

	//请求错误 Bad Request
	BizErrorUserIDMissed                    = NewBizError(BizErrorCodeError, "Missing user ID.")
	BizErrorPermissionDenied                = NewBizError(BizErrorCodeError, "Permission denied.")
	BizErrorInvalidRequestParams            = NewBizError(BizErrorCodeError, "Invalid request parameter.")
	BizErrorContentTypeOnlyAllowJson        = NewBizError(BizErrorCodeError, "Unsupported Content-Type. Only allow application/json.")
	BizErrorContentTypeOnlyAllowJsonAndForm = NewBizError(BizErrorCodeError, "Unsupported Content-Type. Only allow application/json or application/x-www-form-urlencoded.")

	//通用服务端错误
	BizErrorServerError          = NewBizError(BizErrorCodeError, "Server side error.")
	BizErrorOperationFailed      = NewBizError(BizErrorCodeError, "Operation failed.")
	BizErrorBodyParameterInvalid = NewBizError(BizErrorCodeError, "Invalid body parameter")
)
