package api

import (
	"encoding/json"
	"haw.com/models"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	listings, err := models.FetchAll()
	if err != nil {
		panic(err)
	}
	if err = json.NewEncoder(w).Encode(listings); err != nil {
		panic(err)
	}
}
