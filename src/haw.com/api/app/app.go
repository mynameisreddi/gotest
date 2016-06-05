package main

import (
	"haw.com/api"
	"log"
	"net/http"
)

func handleRequests() {
	http.HandleFunc("/listing", api.Index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
