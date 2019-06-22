package service

import (
	"../repository"
	// "../entity"
	"../../infrastructure/dao"
)

type MaterialService struct {}
// var foodEntity *entity.Food
var materialRepository repository.Material = repository.Material{}

func (f MaterialService) GetList (page int, pagesize int) []dao.Material {
	list := materialRepository.Fetch(pagesize, (page - 1) * pagesize)

	return list
}

func (f MaterialService) Add(name string, img string) int64 {
	id := materialRepository.Add(name, img)
	return id
}

