package dao

import (
	"database/sql"
	// "time"
	"fmt"
)

type Food struct {
	Id int64 `sql:"id"`
	Name string `sql:"name"`
	Img string `sql:"img"`
}
var tableName string = "food"

func (f Food) Insert(name string, img string) (int64, error) {
	res, err := DB.Exec("insert into " + tableName + " (`name`, `img`) values (?, ?)", name, img)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (f Food) Select(id int64) (Food, error) {
	row := DB.QueryRow("select `id`, `name`, `img` from " + tableName + 
		" where `id` = ? and `is_delete` = 0", id)
	food := Food{}
	if err := row.Scan(&food.Id, &food.Name, &food.Img); err != nil {
		if err == sql.ErrNoRows {
			return food, nil
		}
		return food, err
	}

	return food, nil
}

func (f *Food) SelectMany(limit int, offset int) ([]Food, error) {
	rows, err := DB.Query("select `id`, `name`, `img` from " + tableName + 
		" where `is_delete` = 0 limit ?, ?", offset, limit)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	foods := []Food{}

	for rows.Next() {
		food := Food{}
		if err := rows.Scan(&food.Id, &food.Name, &food.Img); err != nil {
			return nil, err
		}
		foods = append(foods, food)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	fmt.Println(foods)

	return foods, nil
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
