package service

import (
	"../repository"
	// "../entity"
	"../../infrastructure/dao"
)

type RecipieService struct {}
// var foodEntity *entity.Food
var recipieRepository repository.Recipie

func (f RecipieService) GetList (id int) []dao.Recipie {
	list := recipieRepository.Fetch(id)

	return list
}

func (f FoodService) Add(foodId int64, name string, img string) int64 {
	id := recipieRepository.Add(foodId, name, img)
	return id
}
