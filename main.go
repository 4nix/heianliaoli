package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"./src/userinterface/api"
	"./src/userinterface/controllers"
	"strconv"
	"./src/domain/service"
	"fmt"
)

func main () {
	app := iris.Default()
	app.Get("/api/list", api.List)

	// app.Post("/doEdit", func(ctx iris.Context) {
	// 	id, _ := strconv.ParseInt(ctx.PostValue("Id"), 10, 64)
	// 	foodEntity := entity.Food{}
	// 	if (id > 0) {
	// 		foodEntity.Update(id, ctx.PostValue("Name"), ctx.PostValue("Img"))
	// 	} else {
	// 		foodEntity.Add(ctx.PostValue("Name"), ctx.PostValue("Img"))
	// 	}
		
	// 	ctx.JSON(iris.Map{
	// 		"name": ctx.PostValue("Name"),
	// 		"img": ctx.PostValue("Img"),
	// 	})
	// })

	app.RegisterView(iris.HTML("./src/userinterface/views", ".html"))
	mvc.New(app.Party("/edit")).Handle(new(controllers.EditController))

	mvc.New(app.Party("/list")).Handle(new(controllers.ListController))
	mvc.New(app.Party("/food/edit")).Handle(new(controllers.FoodController))
	mvc.New(app.Party("/material/edit")).Handle(new(controllers.MaterialController))
	mvc.New(app.Party("/recipe/edit")).Handle(new(controllers.RecipeController))


	app.Post("/food/doEdit", func(ctx iris.Context) {
		id, _ := strconv.ParseInt(ctx.PostValue("Id"), 10, 64)
		foodService := service.FoodService{}
		if (id > 0) {
			// foodService.Update(id, ctx.PostValue("Name"), ctx.PostValue("Img"))
		} else {
			foodService.Add(ctx.PostValue("Name"), ctx.PostValue("Img"))
		}
		
		ctx.JSON(iris.Map{
			"name": ctx.PostValue("Name"),
			"img": ctx.PostValue("Img"),
		})
	})

	app.Post("/recipe/doEdit", func(ctx iris.Context) {
		id, _ := strconv.ParseInt(ctx.PostValue("Id"), 10, 64)
		foodId, _ := strconv.ParseInt(ctx.PostValue("foodId"), 10, 64)
		isBest, _ := strconv.Atoi(ctx.PostValue("isBest"))
		// isBest, _ := strconv.ParseInt(ctx.PostValue("Is_best"), 10, 0)
		recipeService := service.RecipeService{}

		fmt.Println(foodId)
		fmt.Println(isBest)

		if (id > 0) {
			// recipeService.Update(id, ctx.PostValue("FoodId"), ctx.PostValue("Name"), ctx.PostValue("Img"), ctx.PostValue("Is_best"))
		} else {
			recipeService.Add(foodId, ctx.PostValue("Info"), isBest)
		}
		
		ctx.JSON(iris.Map{
			"name": ctx.PostValue("Name"),
			"img": ctx.PostValue("Img"),
		})
	})

	app.Post("/material/doEdit", func(ctx iris.Context) {
		id, _ := strconv.ParseInt(ctx.PostValue("Id"), 10, 64)
		materialService := service.MaterialService{}
		if (id > 0) {
			// materialService.Update(id, ctx.PostValue("Name"), ctx.PostValue("Img"))
		} else {
			materialService.Add(ctx.PostValue("Name"), ctx.PostValue("Img"))
		}
		

		ctx.JSON(iris.Map{
			"name": ctx.PostValue("Name"),
			"img": ctx.PostValue("Img"),
		})
	})


	app.Run(iris.Addr(":8082"))
}