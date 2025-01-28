package parser

import (
	"os"

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

func InitConfig() *Config {
	var config Config
	config.init()

	return &config
}
