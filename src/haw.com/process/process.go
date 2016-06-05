package main

import (
	"haw.com/models"
	"haw.com/parsers"
	"log"
	"os"
)

func main() {
	inFile, err := os.Open("/home/reddi/challenge.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer inFile.Close()

	feed, err := parsers.MakeFeed("xxx")
	if err != nil {
		log.Fatal(err)
	}
	listings, err := parsers.Parse(inFile, feed)
	if err != nil {
		log.Fatal(err)
	}
	for i := range listings {
		err = models.Save(listings[i])
		if err != nil {
			log.Fatal(err)
		}
	}
}
