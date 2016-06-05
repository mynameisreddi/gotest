package parsers

import (
    "fmt"
    "errors"
    "strings"
    "strconv"
    "encoding/xml"
    "haw.com/models"
    "haw.com/constants"
)

type XFeed struct {
    XMLName xml.Name `xml:"housing"`
    Items []XFeedItem `xml:"member>items>item"`
}

type XFeedItem struct {
    UniqueObjectID string `xml:"uniqueobjectid"`
    Street string `xml:"street"`
    HouseNumber string `xml:"houseNumber"`
    HouseNumberAddtion string `xml:"houseNumberAddtion"`
    PostalCode string `xml:"postalCode"`
    City string `xml:"City"`
    SubArea string `xml:"SubArea"`
    HouseType int `xml:"HouseType"`
    Projectnaam string `xml:"Projectnaam"`
}

func (item XFeedItem) ConvertToListing() (*models.Listing, error) {
    houseNumber, err := strconv.Atoi(strings.TrimSpace(item.HouseNumber))
    if err != nil {
        return nil, err
    }
    address := models.Address{
        Region: item.SubArea,
        City: item.City,
        Street: item.Street,
        PostalCode: item.PostalCode,
        HouseNumber: houseNumber,
    }
    houseTypesMap := map[int]constants.HouseType {
        0: constants.ROOM,
        1: constants.APPARTEMENT,
        2: constants.GARAGE,
        3: constants.ROOM,
        4: constants.VILLA,
        5: constants.HOUSEBOAT,
        6: constants.STUDIO,
        7: constants.RESIDENCE,
        8: constants.ROOM,
        9: constants.BUNGALOW,
        10: constants.ROOM,
    }
    houseType, found := houseTypesMap[item.HouseType]
    if !found {
        return nil, errors.New(
            fmt.Sprintf("can't find HoueseType %d\n", item.HouseType))
    }
    address.HouseType = houseType

    listing := models.Listing{
        UID: item.UniqueObjectID,
        Address: address,
        Title: item.Projectnaam,
    }
    return &listing, nil
}
