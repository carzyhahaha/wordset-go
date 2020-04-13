package router

import (
	"github.com/kataras/iris/v12"
	"wordset/common"
	"wordset/handler"
)

var wordSetRouter iris.Party

func SetWordSetRouter(app *iris.Application) {
	wordSetRouter := app.Party("/v1")
	wordSetRouter.Post("/wordset_upload", handler.UploadWordSet)
	wordSetRouter.Get("/wordSet", func(ctx iris.Context) {
		response := common.NewApiResponse(0, "test success", nil)
		ctx.Next()
		ctx.JSON(response)
	})

}
