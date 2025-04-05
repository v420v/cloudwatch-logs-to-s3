package config

import "os"

type Config struct {
	Port string
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	return port
}

func LoadConfig() *Config {
	return &Config{
		Port: getPort(),
	}
}
