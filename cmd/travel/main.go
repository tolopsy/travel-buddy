package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"travel-buddy/travel"
)

var LOGFILE = "logs/main.log"

func setLogger(logFiles ...io.Writer) {
	logFiles = append(logFiles, os.Stdout)
	logWriter := io.MultiWriter(logFiles...) // logs error in both console and file
	log.SetOutput(logWriter)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("TravelBuddy: ")
}

func main() {
	// custom log file
	file, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("Error while opening log file:", err)
		return
	}
	defer file.Close()

	setLogger(file)

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
			log.Println(err)
			return
		}

		query.Lng, err = strconv.ParseFloat(req.URL.Query().Get("lng"), 64)
		if err != nil {
			http.Error(responseWriter, err.Error(), http.StatusBadRequest)
			log.Println(err)
			return
		}

		query.Radius, err = strconv.Atoi(req.URL.Query().Get("radius"))
		if err != nil {
			http.Error(responseWriter, err.Error(), http.StatusBadRequest)
			log.Println(err)
			return
		}
		query.CostRangeStr = req.URL.Query().Get("cost")
		places := query.Run()
		respond(responseWriter, req, places)
	}))

	fmt.Println("Listening and Serving at :9000")
	http.ListenAndServe(":9000", http.DefaultServeMux)
}

func respond(writer http.ResponseWriter, req *http.Request, data []interface{}) error {
	publicData := make([]interface{}, len(data))
	for index, value := range data {
		publicData[index] = travel.Public(value)
	}
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "    ")
	return encoder.Encode(publicData)
}

func cors(f http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		f(writer, req)
	}
}
