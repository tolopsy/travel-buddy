package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"travel-buddy/travel"
)

func main() {
	http.HandleFunc("/journeys", func(responseWriter http.ResponseWriter, req *http.Request) {
		respond(responseWriter, req, travel.Journeys)
	})

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
