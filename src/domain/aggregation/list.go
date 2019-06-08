package aggregation

import (
	"../repository"
	// "../entity"
	"../../infrastructure/dao"
)

type FoodList struct {}
// var foodEntity *entity.Food
var foodRepository repository.Food

func (f FoodList) GetList (page int, pagesize int) []dao.Food {
	list := foodRepository.Fetch(pagesize, (page - 1) * pagesize)

	return list
}
