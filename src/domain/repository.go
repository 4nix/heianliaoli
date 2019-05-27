package domain

import (
	"../infrastructure/dao"
)

func GetList () []dao.Food {
	foods, _ := dao.Fetch(0, 200)

	return foods
}