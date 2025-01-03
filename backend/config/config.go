package config

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB


func InitDB() {
	var err error
	Db, err = sql.Open("sqlite3", "../sqlite.db")
	if err != nil {
		panic(err)
	}
}
