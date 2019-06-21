package dao

import (
	"database/sql"
	// "time"
	// "fmt"
)

type Material struct {
	Id int64 `sql:"id"`
	Name string `sql:"name"`
	Img string `sql:"img"`
}

func (material Material) Add(name string, img string) (int64, error) {
	res, err := DB.Exec("insert into `material` (`name`, `img`) values (?, ?)", name, img)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (f Material) Select(id int64) (Material, error) {
	row := DB.QueryRow("select `id`, `name`, `img` from material where `id` = ? and `is_delete` = 0", id)
	material := Material{}
	if err := row.Scan(&material.Id, &material.Name, &material.Img); err != nil {
		if err == sql.ErrNoRows {
			return material, nil
		}
		return material, err
	}

	return material, nil
}

// func Fetch(offset int, limit int) ([]Food, error) {
// 	rows, err := DB.Query("select `id`, `name`, `img` from `food` where `is_delete` = 0 limit ?, ?", offset, limit)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	foods := []Food{}
// 	for rows.Next() {
// 		food := Food{}

// 		if err := rows.Scan(&food.Id, &food.Name, &food.Img); err != nil {
// 			return nil, err
// 		}

// 		fmt.Println(food.Id)
// 		foods = append(foods, food)
// 	}
// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return foods, nil
// }



// func UpdateName(food *Food) error {
// 	_, err := DB.Exec("update `food` set `name` = ?, `u_time` = ? where id = ?", food.Name, time.Now(), food.Id)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func UpdateImg(food *Food) error {
// 	_, err := DB.Exec("update `food` set `img` = ?, `u_time` = ? where id = ?", food.Name, time.Now(), food.Img)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func Delete(food *Food) error {
// 	_, err := DB.Exec("update `food` set `is_delete` = 1, `u_time` = ? where id = ?", time.Now(), food.Id)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
