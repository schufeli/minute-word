package configuration

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

const URL = "https://dwds.de/random"

var api_key = ""

func LoadEnvironment(filename string) {
	godotenv.Load(filename)
}

func LoadApiKey() error {
	token, exists := os.LookupEnv("API_KEY")
	if !exists {
		return errors.New("Error environment variable API_KEY is not set")
	}

	if len(token) == 0 {
		return errors.New("Error environment variable API_KEY is empty")
	}

	api_key = token

	return nil
}

func GetApiKey() string {
	return api_key
}
