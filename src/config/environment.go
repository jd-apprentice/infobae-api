package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	App struct {
		Port string
	}
}

func GetConfig() Config {

	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	return Config{
		App: struct {
			Port string
		}{
			Port: port,
		},
	}
}
