package app

import (
	"dwd-api/pkg/errcode"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Ctx *gin.Context
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

func (r Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, gin.H{
		"data":   data,
		"errno":  0,
		"errmsg": "success",
	})
}

func (r Response) ToResponseList(list interface{}, totalNum int) {
	r.Ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"list":      list,
			"pageNum":   GetPageNum(r.Ctx),
			"pageLimit": GetPageLimit(r.Ctx),
			"totalNum":  totalNum,
		},
		"errno":  0,
		"errmsg": "success",
	})
}

func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{
		"data":   gin.H{},
		"errno":  err.GetErrorCode(),
		"errmsg": err.GetErrorMsg(),
	}
	details := err.GetDetails()
	if len(details) > 0 {
		response["details"] = details
	}
	r.Ctx.JSON(err.StatusCode(), response)
}
