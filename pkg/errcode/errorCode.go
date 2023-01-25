package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	code    int      `json:"code"`
	msg     string   `json:msg`
	details []string `json:"details"`
}

var codes = map[int]string{}

//验证这个错误码是否存在
func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}
	codes[code] = msg
	return &Error{code: code, msg: msg}
}

//返回错误码信息
func (e *Error) GetErrorInfo() string {
	return fmt.Sprintf("错误码:%d, 错误信息:%s", e.code, e.msg)
}

//返回错误码code
func (e *Error) GetErrorCode() int {
	return e.code
}

//返回错误码msg
func (e *Error) GetErrorMsg() string {
	return e.msg
}

//返回错误详情
func (e *Error) GetDetails() []string {
	return e.details
}

//返回错误码信息格式化输出
func (e *Error) GetErrorMsgSprintf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}

func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.details = []string{}
	for _, d := range details {
		newError.details = append(newError.details, d)
	}
	return &newError
}

func (e *Error) StatusCode() int {
	switch e.GetErrorCode() {
	case Success.GetErrorCode():
		return http.StatusOK
	case ServerError.GetErrorCode():
		return http.StatusInternalServerError
	case InvalidParams.GetErrorCode():
		return http.StatusBadRequest
	}

	return http.StatusInternalServerError
}
