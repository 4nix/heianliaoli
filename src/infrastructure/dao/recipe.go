package dao

import (
	"database/sql"
	"fmt"
)

type Recipe struct {
	Id int64 `sql:"id"`
	Food_id int64 `sql:"food_id"`
	Is_best int `sql:"is_best"`
	Info string `sql:"info"`
}

func (f Recipe) Insert(id int64, info string, is_best int) (int64, error) {
	res, err := DB.Exec("insert into recipe (`food_id`, `info`, `is_best`) values (?, ?, ?)", id, info, is_best)
	if err != nil {
		fmt.Println(id)
		fmt.Println(info)
		fmt.Println(is_best)
		fmt.Println(err)
		return 0, err
	}

	return res.LastInsertId()
}

func (f Recipe) Select(id int64) (Recipe, error) {
	row := DB.QueryRow("select `id`, `food_id`, `info`, `is_best` from recipe where `id` = ? and `is_delete` = 0", id)
	recipe := Recipe{}
	if err := row.Scan(&recipe.Id, &recipe.Food_id, &recipe.Info, &recipe.Is_best); err != nil {
		if err == sql.ErrNoRows {
			return recipe, nil
		}
		return recipe, err
	}

	return recipe, nil
}

func (f *Recipe) SelectMany(id int64, limit int, offset int) ([]Recipe, error) {
	rows, err := DB.Query("select `id`, `name`, `img` from recipe where `food_id` = ? and `is_delete` = 0 limit ?, ?", id, offset, limit)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	recipes := []Recipe{}

	for rows.Next() {
		recipe := Recipe{}
		if err := rows.Scan(&recipe.Id, &recipe.Food_id, &recipe.Info, &recipe.Is_best); err != nil {
			return nil, err
		}
		recipes = append(recipes, recipe)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return recipes, nil
}
