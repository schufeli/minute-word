package server

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/schufeli/minute-word/internal/configuration"
	"github.com/schufeli/minute-word/internal/random"
	"github.com/sirupsen/logrus"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	key := r.Header.Get("X-API-KEY")

	if key != configuration.GetApiKey() {
		logrus.Warnf("address: %s made a request with invalid X-API-KEY", r.RemoteAddr)
		http.Error(w, "Invalid API key", http.StatusForbidden)
		return
	}

	logrus.Infof("Request made by address: %s", r.RemoteAddr)

	word, err := random.GetRandomWord(configuration.URL)
	if err != nil {
		logrus.Error(err)
	}

	response, _ := json.Marshal(word)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func ApiKeyCheck(actual string, expected string) bool {
	return strings.Compare(actual, expected) == 0
}
