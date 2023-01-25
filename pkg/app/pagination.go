package app

import (
	"dwd-api/global"
	"dwd-api/pkg/convert"

	"github.com/gin-gonic/gin"
)

func GetPageNum(c *gin.Context) int {
	pageNum := convert.StrTo(c.Query("pageNum")).MustInt()
	if pageNum <= 0 {
		return 1
	}
	return pageNum
}

func GetPageLimit(c *gin.Context) int {
	pageLimit := convert.StrTo(c.Query("pageLimit")).MustInt()
	if pageLimit <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	if pageLimit > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}
	return pageLimit
}

func GetPageOffset(pageNum, pageLimit int) int {
	result := 0
	if pageNum > 0 {
		result = (pageNum - 1) * pageLimit
	}
	return result
}
