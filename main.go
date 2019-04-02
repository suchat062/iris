package main

import (
	"github.com/kataras/iris"
)

func main()  {
	app := iris.New()

	app.Logger().SetLevel("debug")

	views := iris.HTML("./web/views", ".html").
		Layout("layout.html").
		Reload(true)

	app.RegisterView(views)
	app.StaticWeb("/public", "./web/public")

	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("Message", "Golang Socmed")
		ctx.View("index.html")
	})

	app.Run(iris.Addr(":3000"))
}
