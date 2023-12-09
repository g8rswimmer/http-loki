package main

import (
	"log"

	"github.com/g8rswimmer/http-loki/cmd/http-loki/internal/httpx"
	"github.com/g8rswimmer/http-loki/internal/app"
	"github.com/g8rswimmer/http-loki/internal/config"
	"github.com/g8rswimmer/http-loki/internal/mock"
)

func main() {
	value, err := config.LoadValues()
	if err != nil {
		log.Panicf("config values: %v", err)
	}

	handler := httpx.NewRouter()

	if err := mock.AddRoutesFromDirectory(value.MockDir, handler); err != nil {
		log.Panicf("unable to load mock files: %v", err)
	}

	if err := app.Run(value.Port, handler); err != nil {
		log.Panicf("app server: %v", err)
	}
}
