package dao

import (
	"database/sql"
)

var DB *sql.DB

func init () {
	dbconfig := Dbconfig {
		username: "root",
		password: "12345678",
		database: "darkfood",
	}

	DB = connector(dbconfig)
}
