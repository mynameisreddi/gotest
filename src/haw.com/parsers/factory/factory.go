package factory

import (
    "fmt"
    "errors"
    "haw.com/parsers"
)

func MakeFeed(partner string) (interface{}, error) {
    switch partner {
    case "x":
        return &(parsers.XFeed{}), nil
    default:
        return nil, errors.New(
            fmt.Sprintf("can't find Feed for partner %s", partner))
    }
}
