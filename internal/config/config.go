package config

import (
	"log"
	"os"
)

type Config struct {
	HostPort string
}

func LoadConfig() *Config {
	return &Config{
		HostPort: EnvString("HOST_PORT", ":8080"),
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
