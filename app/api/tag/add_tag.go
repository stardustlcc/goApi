package tag

import (
	"dwd-api/pkg/util"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *Hander) AddTag(ctx *gin.Context) {
	defer util.RecoverErr(ctx)
	fmt.Println("hello")
}
