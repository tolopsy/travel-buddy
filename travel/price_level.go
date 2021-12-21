package travel

import (
	"errors"
	"strings"
)

type Price int8

const (
	_ Price = iota
	Price1
	Price2
	Price3
	Price4
	Price5

)

var PriceStrings = map[string]Price{
	"$": Price1,
	"$$": Price2,
	"$$$": Price3,
	"$$$$": Price4,
	"$$$$$": Price5,
}

func (cost Price) String() string {
	for key, value := range PriceStrings {
		if value == cost {
			return key
		}
	}

	return "Invalid"
}

func ParsePrice(s string) Price {
	return PriceStrings[s]
}

type PriceRange struct {
	From Price
	To Price
}

func (priceRange PriceRange) String() string {
	return priceRange.From.String() + "..." + priceRange.To.String()
}

func ParsePriceRange(s string) (PriceRange, error) {
	var priceRange PriceRange
	rangeComponents := strings.Split(s, "...")
	if len(rangeComponents) != 2 {
		return priceRange, errors.New("Invalid Cost Range")
	}

	priceRange.From = ParsePrice(rangeComponents[0])
	priceRange.To = ParsePrice(rangeComponents[1])

	return priceRange, nil
}