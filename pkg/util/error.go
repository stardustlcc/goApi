package util

import (
	"dwd-api/global"
	"dwd-api/pkg/app"
	"dwd-api/pkg/errcode"
	"fmt"

	"github.com/gin-gonic/gin"
)

//捕捉错误
func RecoverErr(ctx *gin.Context) {
	if err := recover(); err != nil {
		msg := fmt.Sprintf("%+v", err)
		global.Logger.Error("err msg: %s", msg)
		response := app.NewResponse(ctx)
		response.ToErrorResponse(errcode.SystemError)
	}
}

//检测错误
func CheckErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
