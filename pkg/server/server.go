package server

import (
	"encoding/json"
	"log"
	"net/http"

	openweather "github.com/medvedev-v/weather-api/pkg/openweather"
	parser "github.com/medvedev-v/weather-api/pkg/parser"
)

type RequestWeather struct {
	Location string `json:"location"`
}

func handleRequest(writer http.ResponseWriter, request *http.Request) {
	config := parser.InitConfig()

	if request.Method != http.MethodGet {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestWeather RequestWeather
	if error := json.NewDecoder(request.Body).Decode(&requestWeather); error != nil {
		http.Error(writer, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if requestWeather.Location == "" {
		http.Error(writer, "Location is required", http.StatusBadRequest)
		return
	}

	response := openweather.AskCurrentWeatherShort(requestWeather.Location, config.OpenweatherKey)

	writer.Header().Set("Content-Type", "application/json")

	if error := json.NewEncoder(writer).Encode(response); error != nil {
		http.Error(writer, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func StartAndServe() {
	http.HandleFunc("/ask", handleRequest)

	log.Println("Server is running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
