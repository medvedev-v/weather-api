package main

import (
	"fmt"
	"os"

	openweather "github.com/medvedev-v/weather-api/pkg/openweather"
	"gopkg.in/yaml.v3"
)

type Config struct {
	OpenweatherKey string `yaml:"openweatherkey"`
}

func (config *Config) init() *Config {
	yamlFile, error := os.ReadFile("config.yaml")
	if error != nil {
		panic(error)
	}
	error = yaml.Unmarshal(yamlFile, config)
	if error != nil {
		panic(error)
	}

	return config
}

func main() {
	var config Config
	config.init()
	fmt.Println(openweather.AskCurrentWeatherShort("Vladivostok", config.OpenweatherKey))
}
