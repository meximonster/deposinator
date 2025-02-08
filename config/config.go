package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	HTTP_PORT     string `yaml:"HTTP_PORT"`
	POSTGRES_HOST string `yaml:"POSTGRES_HOST"`
	POSTGRES_USER string `yaml:"POSTGRES_USER"`
	POSTGRES_PASS string `yaml:"POSTGRES_PASS"`
	STORE_KEY     string `yaml:"STORE_KEY"`
}

func Load() (*Config, error) {
	c := &Config{}
	ymlFile, err := os.ReadFile(".env")
	if err != nil {
		return &Config{}, err
	}
	err = yaml.Unmarshal(ymlFile, c)
	if err != nil {
		return &Config{}, err
	}
	return c, nil
}
