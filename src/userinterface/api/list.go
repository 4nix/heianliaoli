package api

import (
	"../../domain"
	"github.com/kataras/iris"
	"fmt"
)

func List (ctx iris.Context) {
	list := domain.GetList()
	fmt.Println(list)

	for _, value := range list {
		fmt.Println(value.Name)
	}

	ctx.JSON(iris.Map{
		"code": "100",
		"msg": "index",
		"data": list,
	})
}