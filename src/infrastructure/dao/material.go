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

func (material Material) Insert(name string, img string) (int64, error) {
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

func (f *Material) SelectMany(limit int, offset int) ([]Material, error) {
	rows, err := DB.Query("select `id`, `name`, `img` from material where `is_delete` = 0 limit ?, ?", offset, limit)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	materials := []Material{}

	for rows.Next() {
		material := Material{}
		if err := rows.Scan(&material.Id, &material.Name, &material.Img); err != nil {
			return nil, err
		}
		materials = append(materials, material)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return materials, nil
}
