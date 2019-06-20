package service

import (
	"../repository"
	// "../entity"
	"../../infrastructure/dao"
)

type FoodService struct {}
// var foodEntity *entity.Food
var foodRepository repository.Food

func (f FoodService) GetList (page int, pagesize int) []dao.Food {
	list := foodRepository.Fetch(pagesize, (page - 1) * pagesize)

	return list
}

func (f FoodService) Add(name string, img string) int64 {
	id := foodRepository.Add(name, img)
	return id
}
