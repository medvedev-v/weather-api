package openweather

import (
	client "github.com/medvedev-v/weather-api/pkg/client"
	parser "github.com/medvedev-v/weather-api/pkg/parser"
)

type ShortWeather struct {
	Location    string
	Temperature float64
	WeatherType string
	Pressure    int
	Humidity    int
	WindSpeed   float64
}

func AskCurrentWeatherShort(location string, apikey string) ShortWeather {
	var weatherUrl string = "https://api.openweathermap.org/data/2.5/weather?units=metric&q=" + location + "&appid=" + apikey
	response := client.Get(weatherUrl)
	data := parser.JsonToWeatherStruct(response)

	shortResponse := ShortWeather{
		Location:    data.Name,
		Temperature: data.Main.Temp,
		WeatherType: data.Weather[0].Main,
		Pressure:    data.Main.Pressure,
		Humidity:    data.Main.Humidity,
		WindSpeed:   data.Wind.Speed,
	}

	return shortResponse
}
