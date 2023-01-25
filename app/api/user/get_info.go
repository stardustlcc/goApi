package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Hander) GetInfo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name": "hello world",
	})
}
