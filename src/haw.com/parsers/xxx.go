package parsers

import (
	"encoding/xml"
	"errors"
	"fmt"
	"haw.com/constants"
	"haw.com/models"
	"strconv"
	"strings"
)

type XFeed struct {
	XMLName xml.Name    `xml:"housing"`
	Items   []XFeedItem `xml:"member>items>item"`
}

type XFeedItem struct {
	UniqueObjectID     string `xml:"uniqueobjectid"`
	Minprice           string `xml:"minprice"`
	Street             string `xml:"street"`
	HouseNumber        string `xml:"houseNumber"`
	HouseNumberAddtion string `xml:"houseNumberAddtion"`
	PostalCode         string `xml:"postalCode"`
	City               string `xml:"City"`
	SubArea            string `xml:"SubArea"`
	HouseType          int    `xml:"HouseType"`
	Projectnaam        string `xml:"Projectnaam"`
}

func (item XFeedItem) ConvertToListing() (*models.Listing, error) {
	address := models.Address{
		Region:      item.SubArea,
		City:        item.City,
		Street:      item.Street,
		PostalCode:  item.PostalCode,
		HouseNumber: item.HouseNumber,
	}
	houseTypesMap := map[int]constants.HouseType{
		0:  constants.ROOM,
		1:  constants.APPARTEMENT,
		2:  constants.GARAGE,
		3:  constants.ROOM,
		4:  constants.VILLA,
		5:  constants.HOUSEBOAT,
		6:  constants.STUDIO,
		7:  constants.RESIDENCE,
		8:  constants.ROOM,
		9:  constants.BUNGALOW,
		10: constants.ROOM,
	}
	houseType, found := houseTypesMap[item.HouseType]
	if !found {
		return nil, errors.New(
			fmt.Sprintf("can't find HoueseType %d\n", item.HouseType))
	}
	address.HouseType = houseType

	minPrice, err := strconv.ParseInt(strings.TrimSpace(item.Minprice), 10, 64)
	if err != nil {
		return nil, err
	}

	listing := models.Listing{
		UID:      item.UniqueObjectID,
		MinPrice: minPrice * 100,
		Address:  address,
		Title:    item.Projectnaam,
	}
	return &listing, nil
}
