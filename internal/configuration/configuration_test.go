package configuration_test

import (
	"testing"

	"github.com/schufeli/minute-word/internal/configuration"
)

func TestLoadApiKey(t *testing.T) {
	configuration.LoadEnvironment("../../.env.example")

	err := configuration.LoadApiKey()
	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetApiKey(t *testing.T) {
	configuration.LoadEnvironment("../../.env.example")
	configuration.LoadApiKey()

	value := configuration.GetApiKey()

	if len(value) == 0 {
		t.Error("Api key was empty!")
	}
}
