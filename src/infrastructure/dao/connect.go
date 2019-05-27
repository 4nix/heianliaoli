package dao
import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

type Dbconfig struct {
	username	string
	password	string
	database	string
	host		string
}

func connector(dbconfig Dbconfig) *sql.DB {
	db, _ := sql.Open("mysql", dbconfig.username + ":" + dbconfig.password +"@/" + dbconfig.database)

	return db
}
