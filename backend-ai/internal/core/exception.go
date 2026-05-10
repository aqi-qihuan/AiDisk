package core

import "fmt"

// ApiException 业务异常
type ApiException struct {
	Msg  string
	Code int
	Data any
}

func (e *ApiException) Error() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Msg)
}

func NewApiException(msg string, code int) *ApiException {
	return &ApiException{Msg: msg, Code: code}
}

// NewApiExceptionWithData 带数据的业务异常
func NewApiExceptionWithData(msg string, code int, data any) *ApiException {
	return &ApiException{Msg: msg, Code: code, Data: data}
}
