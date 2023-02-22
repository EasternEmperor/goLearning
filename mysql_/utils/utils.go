package utils

import (
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
)

var (
	Db *sql.DB
	err error
)

func init() {
	Db, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/go_test")
	if err != nil {
		panic(err.Error())
	}
}