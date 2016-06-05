package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

//Singleton, representing database pool connections.
var db *sql.DB

func GetDB() *sql.DB {
	if db == nil {
		var err error
		db, err = sql.Open(
			"postgres",
			fmt.Sprintf(
				"user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
				USER, PASSWORD, DB_NAME, HOST, PORT))
		if err != nil {
			panic(err)
		}
	}
	return db
}
