package entity

import(
	"../repository"
)

type Food struct {
	Id			int64
	Name		string
	Img			string
	CTime		string
	UTime		string
	IsDelete	int
}

var foodRepository repository.Food

func (f Food) Add(name string, img string) int64 {
	id := foodRepository.Add(name, img)
	return id
}

func (f Food) Update(id int64, name string, img string) {
	
}

// func (f *Food) GetById(id int64) {

// }
