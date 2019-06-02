package controllers

import (
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris"
)

type EditController struct {}

func (c *EditController) Get() mvc.Result {

	return mvc.View{
		Name: "food/edit.html",
		Data: map[string]interface{} {
			"Title": "Index",
			"Msg": "The first msg",
		},
	}
}

func (c *EditController) GetBy(id int64) mvc.Result {

	return mvc.View{
		Name: "food/edit.html",
		Data: map[string]interface{} {
			"Title": "Index",
			"Msg": "The first msg",
			"Id": id,
		},
	}
}

func (c *EditController) post(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"code": "100",
		"msg":	"success",
	})
}
