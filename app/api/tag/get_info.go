package tag

import (
	"dwd-api/app/models"
	"dwd-api/global"
	"dwd-api/pkg/app"
	"dwd-api/pkg/errcode"
	"dwd-api/pkg/util"

	"github.com/gin-gonic/gin"
)

type TagInfoRequest struct {
	Id int32 `form:"id" binding:"required"`
}

func (h Hander) GetInfo(ctx *gin.Context) {
	defer util.RecoverErr(ctx)
	response := app.NewResponse(ctx)
	tagRequest := &TagInfoRequest{}
	valid, err := app.BindAndValid(ctx, tagRequest)
	if !valid {
		global.Logger.Errorf("app.BindAndValid err:%v", err)
		errRsp := errcode.InvalidParams.WithDetails(err.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}
	tag := models.Tag{Id: tagRequest.Id}
	tags, errs := tag.Info(global.DBEngine)
	if errs != nil {
		global.Logger.Errorf("tag.Info err:%v", errs)
		response.ToErrorResponse(errcode.NotFound)
		return
	}
	response.ToResponse(tags)
}
