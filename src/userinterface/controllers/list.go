package controllers

import (
	"github.com/kataras/iris/mvc"
	// "github.com/kataras/iris"

	"../../domain/aggregation"
)

type ListController struct {}
var listAggregation aggregation.FoodList = aggregation.FoodList{}

func (c *ListController) Get() mvc.Result {
	list := listAggregation.GetList(1, 20)


	// for v, _ := range *list {

	// }

	return mvc.View{
		Name: "list/index.html",
		Data: map[string]interface{} {
			"list": list,
		},
	}
}

// func (c *ListController) GetBy(id int64) mvc.Result {

// 	return mvc.View{
// 		Name: "food/edit.html",
// 		Data: map[string]interface{} {
// 			"Title": "Index",
// 			"Msg": "The first msg",
// 			"Id": id,
// 		},
// 	}
// }

// func (c *ListController) post(ctx iris.Context) {
// 	ctx.JSON(iris.Map{
// 		"code": "100",
// 		"msg":	"success",
// 	})
// }
