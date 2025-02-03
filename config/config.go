package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	POSTGRES_HOST string `yaml:"POSTGRES_HOST"`
	POSTGRES_USER string `yaml:"POSTGRES_USER"`
	POSTGRES_PASS string `yaml:"POSTGRES_PASS"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Load() (*Config, error) {
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
