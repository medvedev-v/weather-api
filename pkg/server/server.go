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

type ResponseWeather struct {
	Location    string  `json:"location"`
	Temperature float64 `json:"temperature"`
	WeatherType string  `json:"weathertype"`
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	config := parser.InitConfig()

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req RequestWeather
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Location == "" {
		http.Error(w, "Name and Value are required", http.StatusBadRequest)
		return
	}

	opeanweatherResponse := openweather.AskCurrentWeatherShort(req.Location, config.OpenweatherKey)

	resp := ResponseWeather{
		Location:    opeanweatherResponse.Location,
		Temperature: opeanweatherResponse.Temperature,
		WeatherType: opeanweatherResponse.WeatherType,
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
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
