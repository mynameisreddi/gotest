package main

import (
	"fmt"
	"haw.com/models"
	"haw.com/parsers"
	"log"
	"os"
	"path/filepath"
)

func main() {
	filename, partner, err := parseArgs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	feed, err := parsers.MakeFeed(partner)
	if err != nil {
		log.Fatal(err)
	}

	listings, err := parsers.Parse(file, feed)
	if err != nil {
		log.Fatal(err)
	}
	for i := range listings {
		if err = models.Save(listings[i]); err != nil {
			log.Fatal(err)
		}
	}
}

func parseArgs() (string, string, error) {
	name := filepath.Base(os.Args[0])
	if len(os.Args) < 3 {
		log.Fatalf("%s required 2 args, passed %d", name, len(os.Args)-1)
	}
	if os.Args[1] == "-h" || os.Args[1] == "--help" {
		err := fmt.Errorf("usage %s infile.xml partner", name)
		return "", "", err
	}

	return os.Args[1], os.Args[2], nil
}
