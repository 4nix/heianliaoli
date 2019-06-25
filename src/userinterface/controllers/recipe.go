package controllers

import (
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris"

	"../../domain/service"
)

type RecipeController struct {}

var materialService service.MaterialService = service.MaterialService{}

func (c *RecipeController) Get() mvc.Result {

	return mvc.View{
		Name: "recipe/edit.html",
		Data: map[string]interface{} {
			"Title": "Index",
			"Msg": "The first msg",
		},
	}
}

func (c *RecipeController) GetBy(foodId int64) mvc.Result {
	materialList := materialService.GetList(1, 1000)
	return mvc.View{
		Name: "recipe/edit.html",
		Data: map[string]interface{} {
			"Title": "Index",
			"materialList": materialList,
			"foodId": foodId,
		},
	}
}

func (c *RecipeController) post(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"code": "100",
		"msg":	"success",
	})
}
