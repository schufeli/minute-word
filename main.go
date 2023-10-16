package main

import (
	"net/http"

	"github.com/schufeli/minute-word/internal/configuration"
	"github.com/schufeli/minute-word/internal/server"
	"github.com/sirupsen/logrus"
)

func main() {
	configuration.LoadEnvironment(".env")

	http.HandleFunc("/", server.Handler)
	http.ListenAndServe(":8000", nil)

	logrus.Info("Server started")
}
