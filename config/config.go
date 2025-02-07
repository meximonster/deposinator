package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	AUTH_TOKEN    string `yaml:"AUTH_TOKEN"`
	POSTGRES_HOST string `yaml:"POSTGRES_HOST"`
	POSTGRES_USER string `yaml:"POSTGRES_USER"`
	POSTGRES_PASS string `yaml:"POSTGRES_PASS"`
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
