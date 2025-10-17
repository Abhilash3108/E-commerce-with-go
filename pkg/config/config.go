package config

import "os"

type Config struct {
	LogLevel      string
}

func LoadConfig() *Config {
	return &Config{
		LogLevel:      getEnv("LOG_LEVEL", "info"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}