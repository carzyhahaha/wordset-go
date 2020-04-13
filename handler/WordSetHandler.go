package handler

import (
	"github.com/kataras/iris/v12"
	"log"
)

func UploadWordSet(ctx iris.Context) {
	body := ctx.Request().Body

	log.Println(body)
}
