package controllers

import (
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris"
)

type RecipeController struct {}

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

	return mvc.View{
		Name: "recipe/edit.html",
		Data: map[string]interface{} {
			"Title": "Index",
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
