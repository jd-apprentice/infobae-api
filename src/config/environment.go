package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	App struct {
		Port string
	}
	Proxies []string
}

func GetConfig() Config {

	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	allowed_proxies := strings.Split(os.Getenv("ALLOWED_PROXIES"), ",")

	return Config{
		App: struct {
			Port string
		}{
			Port: port,
		},
		Proxies: allowed_proxies,
	}
}
