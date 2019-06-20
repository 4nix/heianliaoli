package dao

import (
	"database/sql"
)

type Recipie struct {
	Id int64 `sql:"id"`
	Food_id int64 `sql:"food_id"`
	Is_best int `sql:"is_best"`
	Info string `sql:"info"`
}

func (f Recipie) Insert(id int64, info string, is_best int) (int64, error) {
	res, err := DB.Exec("insert into recipie (`food_id`, `info`, `is_best`) values (?, ?)", id, info, is_best)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (f Recipie) Select(id int64) (Recipie, error) {
	row := DB.QueryRow("select `id`, `food_id`, `info`, `is_best` from recipie where `id` = ? and `is_delete` = 0", id)
	recipie := Recipie{}
	if err := row.Scan(&recipie.Id, &recipie.Food_id, &recipie.Info, &recipie.Is_best); err != nil {
		if err == sql.ErrNoRows {
			return recipie, nil
		}
		return recipie, err
	}

	return recipie, nil
}

func (f *Recipie) SelectMany(limit int, offset int) ([]Recipie, error) {
	rows, err := DB.Query("select `id`, `name`, `img` from recipie where `is_delete` = 0 limit ?, ?", offset, limit)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	recipies := []Recipie{}

	for rows.Next() {
		recipie := Recipie{}
		if err := rows.Scan(&recipie.Id, &recipie.Food_id, &recipie.Info, &recipie.Is_best); err != nil {
			return nil, err
		}
		recipies = append(recipies, recipie)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return recipies, nil
}
