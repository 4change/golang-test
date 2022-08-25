package error_utils

import "fmt"

type XErrorCode int

const (
	XErrorCodeSuccess XErrorCode = iota
	XErrorCodeValuesError
	XErrorCodeConfigError
	XErrorCodeServiceError
	XErrorCodeAggregatorError
	XErrorCodeExpandError
	XErrorCodeCoreError
	XErrorCodeCoreParserError
	XErrorCodeEngineError
	XErrorCodeDataLoaderError
	XErrorCodeDataClusterError
	XErrorCodeRPCError
	XErrorCodeUtilsError

	XErrorCodeEngineFinishedError
)

type XError struct {
	Code        XErrorCode  `json:"code,omitempty"`
	Description string      `json:"desc,omitempty"`
	Message     string      `json:"message,omitempty"`
	ActualError error       `json:"actual_err,omitempty"`
	Track       []string    `json:"track,omitempty"`
	Detail      interface{} `json:"detail,omitempty"`
}

type XErrorServiceDetail struct {
	Endpoint   string
	StatusCode int
}

type XErrorFuncDetail struct {
	FuncName string
	Message  string
}

func newXError(code XErrorCode, desc string) XError {
	return XError{Code: code, Description: desc}
}

// TODO: format XError
func (err XError) Error() string {
	return fmt.Sprintf("code: %d [%s], msg: %s, err: %v, detail: %v", err.Code, err.Description, err.Message, err.ActualError, err.Detail)
}

// declare and success case is nil.
func (err XError) IsNil() bool {
	return err.Code == XErrorCodeSuccess
}

func (err XError) IsNotNil() bool {
	return err.Code != XErrorCodeSuccess
}

func (err XError) WithMessage(msg string) XError {
	err.Message = msg
	return err
}

func (err XError) WithFormatMessage(format string, args ...interface{}) XError {
	err.Message = fmt.Sprintf(format, args...)
	return err
}

func (err XError) WithActualError(e error) XError {
	err.ActualError = e
	return err
}

func (err XError) WithMessageAndError(msg string, e error) XError {
	err.Message = msg
	err.ActualError = e
	return err
}

func (err XError) WithDetail(detail interface{}) XError {
	err.Detail = detail
	return err
}

func (err XError) CatchTrack(t string) {
	err.Track = append(err.Track, t)
}

func (err XError) InitTrack(t string) XError {
	err.Track = append(err.Track, t)
	return err
}

func (err XError) AggregateEvaluateFailed(name string) XError {
	return err.WithDetail(XErrorFuncDetail{
		FuncName: name,
		Message:  "aggregator evaluation error",
	})
}

func (err XError) AggregateVerifyFailed(name string) XError {
	return err.WithDetail(XErrorFuncDetail{
		FuncName: name,
		Message:  "aggregator verify error",
	})
}

func (err XError) ExpandEvaluateFailed(name string) XError {
	return err.WithDetail(XErrorFuncDetail{
		FuncName: name,
		Message:  "expand evaluation error",
	})
}

func (err XError) ExpandVerifyFailed(name string) XError {
	return err.WithDetail(XErrorFuncDetail{
		FuncName: name,
		Message:  "expand verify error",
	})
}

func (err XError) IsEngineFinished() bool {
	return err.Code == XErrorCodeEngineFinishedError
}

var (
	XErrorNil                 = newXError(XErrorCodeSuccess, "")
	XErrorSuccess             = newXError(XErrorCodeSuccess, "success")
	XErrorValuesError         = newXError(XErrorCodeValuesError, "values of constant or feature error")
	XErrorConfigError         = newXError(XErrorCodeConfigError, "config error")
	XErrorServiceError        = newXError(XErrorCodeServiceError, "service failed")
	XErrorAggregatorError     = newXError(XErrorCodeAggregatorError, "aggregator error")
	XErrorExpandError         = newXError(XErrorCodeExpandError, "expand error")
	XErrorCoreError           = newXError(XErrorCodeCoreError, "xfeature core error")
	XErrorCoreParserError     = newXError(XErrorCodeCoreParserError, "xfeature core parser error")
	XErrorEngienError         = newXError(XErrorCodeEngineError, "xfeature engine error")
	XErrorEngineFinishedError = newXError(XErrorCodeEngineFinishedError, "xfeature engine execute finished")
	XErrorDataLoaderError     = newXError(XErrorCodeDataLoaderError, "data loader failed")
	XErrorDataClusterError    = newXError(XErrorCodeDataClusterError, "data cluster operation error")
	XErrorRPCError            = newXError(XErrorCodeRPCError, "xfeature rpc failed")
	XErrorUtilsError          = newXError(XErrorCodeUtilsError, "xfeature utils error")
)
