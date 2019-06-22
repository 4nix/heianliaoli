package controllers

import (
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris"
)

type MaterialController struct {}

func (c *MaterialController) Get() mvc.Result {

	return mvc.View{
		Name: "material/edit.html",
		Data: map[string]interface{} {
			"Title": "Index",
			"Msg": "The first msg",
		},
	}
}

func (c *MaterialController) GetBy(Id int64) mvc.Result {

	return mvc.View{
		Name: "material/edit.html",
		Data: map[string]interface{} {
			"Title": "Index",
			"Msg": "The first msg",
			"Id": Id,
		},
	}
}

func (c *MaterialController) post(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"code": "100",
		"msg":	"success",
	})
}
