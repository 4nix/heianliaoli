package controllers

import (
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris"
)

type FoodController struct {}

func (c *FoodController) Get() mvc.Result {

	return mvc.View{
		Name: "food/edit.html",
		Data: map[string]interface{} {
			"Title": "Index",
			"Msg": "The first msg",
		},
	}
}

func (c *FoodController) GetBy(id int64) mvc.Result {

	return mvc.View{
		Name: "food/edit.html",
		Data: map[string]interface{} {
			"Title": "Index",
			"Msg": "The first msg",
			"Id": id,
		},
	}
}

func (c *FoodController) post(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"code": "100",
		"msg":	"success",
	})
}
