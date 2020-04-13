package router

import (
	"github.com/kataras/iris/v12"
	"wordset/common"
)

var ruleRouter iris.Party

func SetRuleSetRouter(app *iris.Application) {
	ruleRouter := app.Party("/v1")
	ruleRouter.Get("/rule", func(ctx iris.Context) {
		response := common.NewApiResponse(0, "test success", nil)
		ctx.Next()
		ctx.JSON(response)
	})

}
