package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var db *sql.DB

func GetDB() *sql.DB {
	DB_USER, DB_PASSWORD, DB_NAME := "postgres", "", "postgres"
	if db == nil {
		var err error
		db, err = sql.Open(
			"postgres",
			fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
				DB_USER, DB_PASSWORD, DB_NAME))
		if err != nil {
			panic(err)
		}
	}
	return db
}
