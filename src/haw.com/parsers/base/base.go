package base

import (
    "io"
	"reflect"
    "encoding/xml"
    "haw.com/models"
)

type Feed interface {}

type FeedItem interface {
    ConvertToListing () (*models.Listing, error)
}

func Parse(reader io.Reader, feed interface{}) ([]models.Listing, error) {
    if err := xml.NewDecoder(reader).Decode(feed); err != nil {
        return nil, err
    }

    items := reflect.ValueOf(feed).Elem().FieldByName("Items")
    listings := make([]models.Listing, items.Len())

    for i := 0; i < items.Len(); i ++ {
        item := items.Index(i).Interface().(FeedItem)
        listing, err := item.ConvertToListing()
        if err != nil {
            return listings, err
        }
        listings[i] = *listing
    }
	return listings, nil
}
