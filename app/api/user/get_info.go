package user

import (
	"dwd-api/pkg/app"
	client "dwd-api/pkg/http"
	"dwd-api/pkg/util"

	"github.com/gin-gonic/gin"
)

func (h Hander) GetInfo(ctx *gin.Context) {
	defer util.RecoverErr(ctx)
	httpParamsMap := map[string]string{
		"plat": "pc",
		"aid":  "mxzzfmt1681815",
	}
	bodyString := util.BuildBody(httpParamsMap)
	httpRequest := &client.HttpRequest{
		Url:  "https://api.auto.sina.cn/api/track/xy_click.json",
		Body: bodyString,
	}
	httpResponse, _ := client.GetHttpClient(int64(2000)).Post(httpRequest)
	response := app.NewResponse(ctx)
	response.ToResponse(string(httpResponse.Body))
}
