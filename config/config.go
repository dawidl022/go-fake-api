package config

type Config struct {
	DatabaseUrl string
	BaseDir     string
}

func LoadDefaults() *Config {
	return &Config{
		DatabaseUrl: "postgresql://user:password@localhost:5432/fake?sslmode=disable",
		BaseDir:     "",
	}
}
