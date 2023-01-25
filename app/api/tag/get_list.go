package tag

import (
	"dwd-api/app/models"
	"dwd-api/global"
	"dwd-api/pkg/app"
	"dwd-api/pkg/errcode"
	"dwd-api/pkg/util"

	"github.com/gin-gonic/gin"
)

type TagListRequest struct {
	Name      string `form:"name"`
	State     int    `form:"state"`
	PageNum   int32  `form:"pageNum" binding:"required"`
	PageLimit int32  `form:"pageLimit" binding:"required"`
}

func (h *Hander) GetList(ctx *gin.Context) {
	defer util.RecoverErr(ctx)
	response := app.NewResponse(ctx)
	tagRequest := &TagListRequest{}
	valid, err := app.BindAndValid(ctx, tagRequest)
	if !valid {
		global.Logger.Errorf("app.BindAndValid err:%v", err)
		errRsp := errcode.InvalidParams.WithDetails(err.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	tag := models.Tag{Name: tagRequest.Name}
	totalRows, errs := tag.Count(global.DBEngine)
	if errs != nil {
		global.Logger.Errorf("tag.Count err:%v", err)
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}

	pageOffset := app.GetPageOffset(app.GetPageNum(ctx), app.GetPageLimit(ctx))
	tags, errs := tag.List(global.DBEngine, pageOffset, app.GetPageLimit(ctx))
	if errs != nil {
		global.Logger.Errorf("tag.List err:%v", err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}

	response.ToResponseList(tags, totalRows)
}
