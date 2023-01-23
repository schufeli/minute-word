package random_test

import (
	"testing"

	"github.com/schufeli/minute-word/internal/configuration"
	"github.com/schufeli/minute-word/internal/random"
)

func TestGetRandomWord(t *testing.T) {
	_, err := random.GetRandomWord(configuration.URL)
	if err != nil {
		t.Errorf(err.Error())
	}
}
