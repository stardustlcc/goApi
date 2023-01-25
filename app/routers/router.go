package routers

import (
	"dwd-api/app/api/tag"
	"dwd-api/app/api/user"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	tag := tag.New()
	tagApi := r.Group("/tag")
	{
		tagApi.GET("/info", tag.GetInfo)
		tagApi.GET("/list", tag.GetList)
	}

	user := user.New()
	userApi := r.Group("/user")
	{
		userApi.GET("/info", user.GetInfo)
	}
	return r
}
