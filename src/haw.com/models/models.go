package models

import (
	"database/sql"
	"haw.com/constants"
)

type Address struct {
	Country     string
	Region      string
	City        string
	Street      string
	PostalCode  string
	SubArea     string
	HouseNumber int
	constants.HouseType
}

type Listing struct {
	UID      string
	MinPrice int64
	Title    string
	Address
}

func ScanListing(rows *sql.Rows) (*Listing, error) {
	var listing Listing
	var houseType string

	err := rows.Scan(
		&listing.UID,
		&listing.MinPrice,
		&listing.Title,
		&listing.Address.Country,
		&listing.Address.Region,
		&listing.Address.City,
		&listing.Address.Street,
		&listing.Address.PostalCode,
		&listing.Address.SubArea,
		&listing.Address.HouseNumber,
		&houseType,
	)
	listing.Address.HouseType = constants.HouseType(houseType)
	if err != nil {
		return nil, err
	}
	return &listing, nil
}

func FetchOne(UID string, db *sql.DB) (*Listing, error) {
	rows, err := db.Query(
		`SELECT uid, minprice, title, country, region, city, street,
                postalcode, subarea, housenumber, housetype
         FROM listing WHERE uid = $1`, UID)

	defer rows.Close()

	if err != nil {
		return nil, err
	}
	if !rows.Next() {
		return nil, nil
	}
	listing, err := ScanListing(rows)
	if err != nil {
		return nil, err
	}
	return listing, nil
}

func FetchAll(db *sql.DB) ([]Listing, error) {
	rows, err := db.Query(
		`SELECT uid, minprice, title, country, region, city, street,
                postalcode, subarea, housenumber, housetype
         FROM listing`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var listings []Listing
	for rows.Next() {
		listing, err := ScanListing(rows)
		if err != nil {
			return listings, err
		}
		listings = append(listings, *listing)
	}

	return listings, nil
}

func Save(listing Listing, db *sql.DB) error {
	if l, _ := FetchOne(listing.UID, db); l != nil {
		_, err := db.Exec(
			"UPDATE listing set title = $1, minprice = $2, country = $3, "+
				"region = $4, city = $5, street = $6, postalcode = $7, "+
				"subarea = $8, housenumber = $9, housetype = $10 WHERE uid = $11",
			listing.Title,
			listing.MinPrice,
			listing.Address.Country,
			listing.Address.Region,
			listing.Address.City,
			listing.Address.Street,
			listing.Address.PostalCode,
			listing.Address.SubArea,
			listing.Address.HouseNumber,
			string(listing.Address.HouseType),
			listing.UID,
		)
		if err != nil {
			return err
		}
		return nil
	}
	_, err := db.Exec(
		"INSERT INTO listing (uid, minprice, title, country, region, "+
			"city, street, postalcode, subarea, "+
			"housenumber, housetype) VALUES ($1, $2, $3, "+
			"$4, $5, $6, $7, $8, $9, $10, $11)",
		listing.UID,
		listing.MinPrice,
		listing.Title,
		listing.Address.Country,
		listing.Address.Region,
		listing.Address.City,
		listing.Address.Street,
		listing.Address.PostalCode,
		listing.Address.SubArea,
		listing.Address.HouseNumber,
		string(listing.Address.HouseType),
	)
	if err != nil {
		return err
	}
	return nil
}
