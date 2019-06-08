package repository

import (
	"../../infrastructure/dao"
)

type Food struct {}

var daoFood dao.Food = dao.Food{}

func (food Food) Add (name string, img string) (int64) {
	id, _ := daoFood.Insert(name, img)

	return id
}

func (f Food) FetchOne (id int64) dao.Food {
	food, _ := daoFood.Select(id)

	return food
}

func (f Food) Fetch (limit int, offset int) []dao.Food {
	foods, _ := daoFood.SelectMany(limit, offset)

	return foods
}
