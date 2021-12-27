package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"travel-buddy/travel"	
)

func main() {
	
	http.HandleFunc("/journeys", cors(func(responseWriter http.ResponseWriter, req *http.Request) {
		respond(responseWriter, req, travel.Journeys)
	}))

	http.HandleFunc("/recommendations", cors(func(responseWriter http.ResponseWriter, req *http.Request) {
		query := &travel.Query{
			Journey: strings.Split(req.URL.Query().Get("journey"), "|"),
		}
		var err error

		query.Lat, err = strconv.ParseFloat(req.URL.Query().Get("lat"), 64)
		if err != nil {
			http.Error(responseWriter, err.Error(), http.StatusBadRequest)
			return
		}

		query.Lng, err = strconv.ParseFloat(req.URL.Query().Get("lng"), 64)
		if err != nil {
			http.Error(responseWriter, err.Error(), http.StatusBadRequest)
			return
		}

		query.Radius, err = strconv.Atoi(req.URL.Query().Get("radius"))
		if err != nil {
			http.Error(responseWriter, err.Error(), http.StatusBadRequest)
			return
		}
		query.CostRangeStr = req.URL.Query().Get("cost")
		places := query.Run()
		respond(responseWriter, req, places)
	}))

	fmt.Println("Listening and Serving at :8080")
	http.ListenAndServe(":8080", http.DefaultServeMux)
}

func respond(writer http.ResponseWriter, req *http.Request, data []interface{}) error {
	publicData := make([]interface{}, len(data))
	for index, value := range data {
		publicData[index] = travel.Public(value)
	}
	return json.NewEncoder(writer).Encode(publicData)
}

func cors(f http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		f(writer, req)
	}
}
