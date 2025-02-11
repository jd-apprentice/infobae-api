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

	if len(allowed_proxies) == 0 {
		allowed_proxies = []string{"127.0.0.1", "noticias.jonathan.com.ar"}
	}

	return Config{
		App: struct {
			Port string
		}{
			Port: port,
		},
		Proxies: allowed_proxies,
	}
}
