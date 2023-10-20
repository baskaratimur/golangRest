package database

import (
	"database/sql"
	"project_baskara/helper"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/golang-database")
	helper.Panicerr(err)
	return db

}
