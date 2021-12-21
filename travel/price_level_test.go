package travel_test

import (
	"testing"
	"travel-buddy/travel"

	"github.com/cheekybits/is"
)

func TestPriceValues(t *testing.T) {
	is := is.New(t)

	is.Equal(int(travel.Price1), 1)
	is.Equal(int(travel.Price2), 2)
	is.Equal(int(travel.Price3), 3)
	is.Equal(int(travel.Price4), 4)
	is.Equal(int(travel.Price5), 5)
}

func TestPriceStrings(t *testing.T) {
	is := is.New(t)

	is.Equal(travel.Price1.String(), "$")
	is.Equal(travel.Price2.String(), "$$")
	is.Equal(travel.Price3.String(), "$$$")
	is.Equal(travel.Price4.String(), "$$$$")
	is.Equal(travel.Price5.String(), "$$$$$")
}

func TestParsePrice(t *testing.T) {
	is := is.New(t)

	is.Equal(travel.Price1, travel.ParsePrice("$"))
	is.Equal(travel.Price2, travel.ParsePrice("$$"))
	is.Equal(travel.Price3, travel.ParsePrice("$$$"))
	is.Equal(travel.Price4, travel.ParsePrice("$$$$"))
	is.Equal(travel.Price5, travel.ParsePrice("$$$$$"))
}

func TestPriceRangeString(t *testing.T) {
	is := is.New(t)
	pRange := travel.PriceRange{
		From: travel.Price1,
		To:   travel.Price4,
	}

	is.Equal(pRange.String(), "$...$$$$")
}

func TestParsePriceRange(t *testing.T) {
	is := is.New(t)
	pRange, err := travel.ParsePriceRange("$$...$$$")
	is.NoErr(err)
	is.Equal(pRange.From, travel.Price2)
	is.Equal(pRange.To, travel.Price3)

	pRange, err = travel.ParsePriceRange("$...$$$$$")
	is.NoErr(err)
	is.Equal(pRange.From, travel.Price1)
	is.Equal(pRange.To, travel.Price5)

	pRange, err = travel.ParsePriceRange("$$...$$$...$$$$")
	is.Err(err)
	is.Equal(int(pRange.From), 0)
	is.Equal(int(pRange.To), 0)
}
