package config

import (
	"log"
	"os"
)

type Config struct {
	Host string
}

func LoadConfig() *Config {
	return &Config{
		Host: EnvString("HOST_ADDR", "0.0.0.0") + EnvString("HOST_PORT", ":8080"),
	}
}

func EnvString(key, defaultValue string) string {
	val := os.Getenv(key)
	log.Println(val)
	if val == "" {
		return defaultValue
	}

	return val
}
