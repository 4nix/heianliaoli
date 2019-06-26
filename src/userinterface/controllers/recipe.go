package controllers

import (
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris"
	"sync"
	"fmt"

	"../../domain/service"
	"../../infrastructure/dao"
)

type RecipeController struct {}

var materialService service.MaterialService = service.MaterialService{}
var recipeService service.RecipeService = service.RecipeService{}

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
	

	materialList := make(chan []dao.Material)
	recipeList := make(chan []dao.Recipe)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		fmt.Println("fbbb")
		materialList <- materialService.GetList(1, 1000)
	}()

	go func() {
		defer wg.Done()
		recipeList <- recipeService.GetList(foodId, 1, 1000)
	}()

	go func() {
		fmt.Println("wait")
		wg.Wait()
		close(materialList)
		close(recipeList)
	}()
	fmt.Println("end")
	

	return mvc.View{
		Name: "recipe/edit.html",
		Data: map[string]interface{} {
			"Title": "Index",
			"materialList": materialList,
			"recipeList": recipeList,
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
