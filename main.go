package main

import (
	"fmt"

	openweather "github.com/medvedev-v/weather-api/pkg/openweather"
)

const weathersource string = "openweather"
const apikey string = "openweatherkey"

func main() {
	fmt.Println(openweather.AskForecast("Vladivostok", apikey))
}
