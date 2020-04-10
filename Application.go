package main

import (
	"encoding/json"
	"github.com/kataras/iris/v12"
	"log"
	"strconv"
	"wordset/common"
)

func main() {
	app := iris.New()
	app.Use(before)
	app.Done(after)
	app.RegisterView(iris.HTML("./views", ".html"))
	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("message", "hello iris")
		ctx.View("hello.html")
	})
	app.Get("/user/{id:uint64}", func(ctx iris.Context) {
		response := common.NewApiResponse(0, "1", nil)

		rs, _ := json.Marshal(response)
		ctx.Write(rs)
		ctx.Next()
	})

	app.Post("/testpost", func(ctx iris.Context) {
		ctx.Writef("abc")

	})

	app.Macros().Get("int").RegisterFunc("min",
		func(minVal int) func(string) bool {
			return func(param string) bool {
				v, err := strconv.Atoi(param)
				if err != nil {
					return false
				}
				return v >= minVal
			}
		})

	app.Get("/min/{v:int min(3)}", func(ctx iris.Context) {
		v, _ := ctx.Params().GetInt("v")
		ctx.Writef("get v: %d", v)
	})

	app.Listen(":9900")

	//t1 := ac_automation.NewTireTree()
	//t1.AddWord("aaaaa")
	//t1.AddWord("aaaab")
	//t1.AddWord("aaaac")
	//t1.FailLink()
	//
	//fmt.Println(t1.Query("aaaabaaaac"))

}

func before(ctx iris.Context) {
	host := ctx.Request().Host
	path := ctx.Path()
	body, err := ctx.GetBody()

	log.Println("host", host, "uri:", path)
	if err != nil {
		log.Println("body:", body)
	}
	ctx.Record()
	ctx.Next()
}

func after(ctx iris.Context) {

	response := ctx.Recorder().Body()
	log.Println(string(response))
	log.Println("finish...")
}
