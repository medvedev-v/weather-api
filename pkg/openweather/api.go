package openweather

import (
	client "github.com/medvedev-v/weather-api/pkg/client"
	parser "github.com/medvedev-v/weather-api/pkg/parser"
)

type ShortWeather struct {
	location    string
	temperature float64
	weatherType string
}

func AskForecast(location string, apikey string) ShortWeather {
	var weatherUrl string = "https://api.openweathermap.org/data/2.5/weather?q=" + location + "&appid=" + apikey
	response := client.Get(weatherUrl)
	data := parser.JsonToWeatherStruct(response)

	shortResponse := ShortWeather{
		location:    data.Name,
		temperature: data.Main.Temp,
		weatherType: data.Weather[0].Description,
	}

	return shortResponse
}
