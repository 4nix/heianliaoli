package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"./src/userinterface/api"
	"./src/userinterface/controllers"
	// "strconv"
)

func main () {
	app := iris.Default()
	app.Get("/api/list", api.List)

	app.Post("/doEdit", func(ctx iris.Context) {
		// id, _ := strconv.ParseInt(ctx.PostValue("Id"), 10, 64)
		// if (id > 0) {
		// 	domain.Update(id, ctx.PostValue("Name"), ctx.PostValue("Img"))
		// } else {
		// 	domain.Add(ctx.PostValue("Name"), ctx.PostValue("Img"))
		// }
		
		ctx.JSON(iris.Map{
			"name": ctx.PostValue("Name"),
			"img": ctx.PostValue("Img"),
		})
	})

	app.RegisterView(iris.HTML("./src/userinterface/views", ".html"))
	mvc.New(app.Party("/edit")).Handle(new(controllers.EditController))
	mvc.New(app.Party("/list")).Handle(new(controllers.ListController))

	app.Run(iris.Addr(":8082"))
}