package service

import (
	"../repository"
	// "../entity"
	"../../infrastructure/dao"
)

type RecipeService struct {}
// var foodEntity *entity.Food
var recipeRepository repository.Recipe = repository.Recipe{}

func (f RecipeService) GetList (id int64, page int, pagesize int) []dao.Recipe {
	list := recipeRepository.Fetch(id, pagesize, (page - 1) * pagesize)

	return list
}

func (f RecipeService) Add(foodId int64, info string, isBest int) int64 {
	id := recipeRepository.Add(foodId, info, isBest)
	return id
}
