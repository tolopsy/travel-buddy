package travel

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"sync"
	"time"
)

var APIKey string

type Place struct {
	*googleGeometry `json:"geometry"`
	Name            string         `json:"name"`
	Icon            string         `json:"icon"`
	Photos          []*googlePhoto `json:"photos"`
	Vicinity        string         `json:"vicinity"`
}

func (p *Place) Public() interface{} {
	return map[string]interface{}{
		"name":     p.Name,
		"icon":     p.Icon,
		"photos":   p.Photos,
		"vicinity": p.Vicinity,
		"lat":      p.Lat,
		"lng":      p.Lng,
	}
}

type googleResponse struct {
	Results []*Place `json:"results"`
}

type googleGeometry struct {
	*googleLocation `json:"location"`
}

type googleLocation struct {
	Lat float64 `json:"Lat"`
	Lng float64 `json:"Lng"`
}

type googlePhoto struct {
	PhotoRef string `json:"photo_reference"`
	URL      string `json:"url"`
}

type Query struct {
	Lat          float64
	Lng          float64
	Journey      []string
	Radius       int
	CostRangeStr string
}

func (q *Query) find(types string) (*googleResponse, error) {
	queryUrl := "https://maps.googleapis.com/maps/api/place/nearbysearch/json"

	queryParams := make(url.Values)
	queryParams.Set("location", fmt.Sprintf("%g,%g", q.Lat, q.Lng))
	queryParams.Set("radius", fmt.Sprintf("%d", q.Radius))
	queryParams.Set("types", types)
	queryParams.Set("key", APIKey)

	if len(q.CostRangeStr) > 0 {
		costRange, err := ParsePriceRange(q.CostRangeStr)

		if err != nil {
			return nil, err
		}

		queryParams.Set("minprice", fmt.Sprintf("%d", int(costRange.From)-1))
		queryParams.Set("maxprice", fmt.Sprintf("%d", int(costRange.To)-1))
	}

	queryResponse, err := http.Get(queryUrl + "?" + queryParams.Encode())
	if err != nil {
		return nil, err
	}

	defer queryResponse.Body.Close()
	var gResponse googleResponse

	if err := json.NewDecoder(queryResponse.Body).Decode(&gResponse); err != nil {
		return nil, err
	}

	return &gResponse, nil
}

func (q *Query) Run() []interface{} {
	rand.Seed(time.Now().UnixNano())
	var waitGroup sync.WaitGroup
	var accessGiver sync.Mutex

	places := make([]interface{}, len(q.Journey))
	for index, placeTypes := range q.Journey {
		waitGroup.Add(1)
		go func(types string, i int) {
			defer waitGroup.Done()
			response, err := q.find(types)
			if err != nil {
				log.Println("Failed to find places:", err)
				return
			}

			if len(response.Results) == 0 {
				log.Println("No places found for:", types)
				return
			}

			for _, result := range response.Results {
				for _, photo := range result.Photos {
					photo.URL = "https://maps.googleapis.com/maps/api/place/photo?" + "maxwidth=1000&photoreference=" +
					photo.PhotoRef + "&key=" + APIKey
				}
			}

			randomizer := rand.Intn(len(response.Results))
			accessGiver.Lock()
			places[i] = response.Results[randomizer]
			accessGiver.Unlock()
		}(placeTypes, index)
	}

	waitGroup.Wait()
	return places
}
