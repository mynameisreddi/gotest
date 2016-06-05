package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"haw.com/models"
	"haw.com/parsers/base"
	"haw.com/parsers/factory"
	"log"
	"os"
)

func main() {
	inFile, err := os.Open("/home/reddi/challenge.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer inFile.Close()

	feed, err := factory.MakeFeed("x")
	if err != nil {
		log.Fatal(err)
	}
	listings, err := base.Parse(inFile, feed)
	if err != nil {
		log.Fatal(err)
	}
	DB_USER, DB_PASSWORD, DB_NAME := "postgres", "", "postgres"
	dbInfo := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME,
	)
	db, err := sql.Open("postgres", dbInfo)
	for i := range listings {
		err = models.Save(listings[i], db)
		if err != nil {
			log.Fatal(err)
		}
	}
}
