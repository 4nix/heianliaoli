package api

import (
	"../../domain/aggregation"
	"github.com/kataras/iris"
	// "fmt"
	// "../../infrastructure/dao"
)

var listAggregation aggregation.FoodList = aggregation.FoodList{}
// var daoa *dao.Food

func List (ctx iris.Context) {
	list := listAggregation.GetList(1, 10)
	// fmt.Println(list)
	// c := dao.Food{}
	// list,_ := c.SelectMany(1, 1)


	// for _, value := range *list {
	// 	fmt.Println(value.Name)
	// }

	ctx.JSON(iris.Map{
		"code": "100",
		"msg": "index",
		"data": list,
	})
}