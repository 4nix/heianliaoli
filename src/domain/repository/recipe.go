package repository

import (
	"../../infrastructure/dao"
)

type Recipe struct {}

var daoRecipe dao.Recipe = dao.Recipe{}

func (f Recipe) Add (foodId int64, info string, isBest int) (int64) {
	id, _ := daoRecipe.Insert(foodId, info, isBest)

	return id
}

func (f Recipe) FetchOne (id int64) dao.Recipe {
	recipe, _ := daoRecipe.Select(id)

	return recipe
}

func (f Recipe) Fetch (id int64, limit int, offset int) []dao.Recipe {
	recipes, _ := daoRecipe.SelectMany(id, limit, offset)

	return recipes
}
