package domain

import (
	"../infrastructure/dao"
)

func GetList () []dao.Food {
	foods, _ := dao.Fetch(0, 200)

	return foods
}

func Add (name string, img string) {
	var food = dao.Food{
		Name: name,
		Img: img,
	}

	dao.Add(&food)
}

func Update (id int64, name string, img string) {
	var food = dao.Food {
		Name: name,
		Img: img,
		Id: id,
	}

	dao.UpdateName(&food)
}

