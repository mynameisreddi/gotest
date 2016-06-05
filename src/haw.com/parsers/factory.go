package parsers

import (
	"errors"
	"fmt"
)

//Factory function for creating concrete Feed value from partner name.
func MakeFeed(partner string) (interface{}, error) {
	switch partner {
	case "xxx":
		return &(XFeed{}), nil
	default:
		return nil, errors.New(
			fmt.Sprintf("can't find Feed for partner %s", partner))
	}
}
