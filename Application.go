package main

import (
	"fmt"
	ac_automation "wordset/util/ac-automation"
)

func main()  {
	//app := iris.New()
	//app.RegisterView(iris.HTML("./views", ".html"))
	//app.Get("/", func(ctx iris.Context) {
	//	ctx.ViewData("message", "hello iris")
	//	ctx.View("hello.html")
	//})
	//app.Get("/user/{id:uint64}", func(ctx iris.Context) {
	//	userID, _ := ctx.Params().GetUint64("id")
	//	ctx.Writef("User ID: %d", userID)
	//})
	//
	//app.Listen(":9900")

	t1 := ac_automation.NewTireTree()
	t1.AddWord("aaaaa")
	t1.AddWord("aaaab")
	t1.AddWord("aaaac")
	t1.FailLink()

	fmt.Println(t1.Query("aaaabaaaac"))


}

