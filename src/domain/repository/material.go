package repository

import (
	"../../infrastructure/dao"
)

type Material struct {}

var daoMaterial dao.Material = dao.Material{}

func (f Material) Add (name string, img string) (int64) {
	id, _ := daoMaterial.Insert(name, img)

	return id
}

func (f Material) FetchOne (id int64) dao.Material {
	material, _ := daoMaterial.Select(id)

	return material
}

func (f Material) Fetch (limit int, offset int) []dao.Material {
	materials, _ := daoMaterial.SelectMany(limit, offset)

	return materials
}
