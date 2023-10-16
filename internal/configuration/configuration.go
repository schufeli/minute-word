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
	err := loadApiKey()
	if err != nil {
		logrus.Fatal(err)
		panic("could not load config",err)
	}
	logrus.Info("Loaded Envrionment")
}

func loadApiKey() error {
	token, exists := os.LookupEnv("API_KEY")
	if !exists {
		return errors.New("Error environment variable API_KEY is not set")
	}

	if len(token) == 0 {
		return errors.New("Error environment variable API_KEY is empty")
	}

	api_key = token
	logrus.Info("Loaded API_KEY Environment variable")
	return nil
}

func GetApiKey() string {
	return api_key
}
