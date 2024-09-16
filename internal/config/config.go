package config

import "os"

type Config struct {
	Port     string
	APIToken string
}

func Load() *Config {
	return &Config{
		Port:     getEnv("PORT", "8645"),
		APIToken: getEnv("API_TOKEN", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
