//Package parsers implements parsing of different types of feeds.
package parsers

import (
	"encoding/xml"
	"haw.com/models"
	"io"
	"reflect"
)

type Feed interface{}

type FeedItem interface {
	//Converts parsed xml data to uniform model Listing.
	ConvertToListing() (*models.Listing, error)
}

//Constructs slice of Listings from feed xml data.
//
//Due to reflection mechanism this function can work with any feed type.
//Use FeedItem.ConvertToListing method for converting every item to Listing.
func Parse(reader io.Reader, feed Feed) ([]models.Listing, error) {
	if err := xml.NewDecoder(reader).Decode(feed); err != nil {
		return nil, err
	}

	items := reflect.ValueOf(feed).Elem().FieldByName("Items")
	listings := make([]models.Listing, items.Len())

	for i := 0; i < items.Len(); i++ {
		item := items.Index(i).Interface().(FeedItem)
		listing, err := item.ConvertToListing()
		if err != nil {
			return listings, err
		}
		listings[i] = *listing
	}
	return listings, nil
}
