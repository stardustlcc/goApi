package errcode

var (
	SystemError   = NewError(-1, "系统错误")
	Success       = NewError(0, "成功")
	ServerError   = NewError(10000000, "服务器内部错误")
	InvalidParams = NewError(10000001, "入参错误")
	NotFound      = NewError(10000002, "找不到")
)
