package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	MONGO_URI   string
	DATABASE    string
	API_VERSION string
	COLLECTION  string
	PORT        int
	GIN_MODE    string
}

func InitAppConfig() *AppConfig {
	error := godotenv.Load()
	if error != nil {
		log.Fatal("Error loading .env file")
	}
	appConfig := AppConfig{

		MONGO_URI:   os.Getenv("MONGO_URI"),
		DATABASE:    os.Getenv("DATABASE"),
		COLLECTION:  os.Getenv("COLLECTION"),
		API_VERSION: os.Getenv("API_VERSION"),
		GIN_MODE:    os.Getenv("GIN_MODE"),
		PORT: func() int {
			var port = os.Getenv("PORT")
			value, error := strconv.Atoi(port)
			if error != nil {
				return value
			}
			return 8080
		}(),
	}
	return &appConfig

}
