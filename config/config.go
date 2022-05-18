package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	DatabaseUrl string `envconfig:"DATABASE_URL" default:"postgresql://user:password@postgres:5432/fake?sslmode=disable"`
	BaseDir     string
}

func LoadDefaults() (*Config, error) {
	var c Config
	err := envconfig.Process("fake", &c)
	return &c, err
}
