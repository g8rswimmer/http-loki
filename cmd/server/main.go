package main

import (
	"log"

	"github.com/g8rswimmer/http-loki/internal/app"
	"github.com/g8rswimmer/http-loki/internal/config"
)

func main() {
	value, err := config.LoadValues()
	if err != nil {
		log.Panicf("config values: %v", err)
	}

	if err := app.Run(value); err != nil {
		log.Panicf("app server: %v", err)
	}
}
