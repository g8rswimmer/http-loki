package mux

import (
	"net/http"

	"github.com/g8rswimmer/http-loki/internal/config"
	"github.com/gorilla/mux"
)

func Handler(values *config.Values) http.Handler {
	router := mux.NewRouter()

	router.Methods(http.MethodGet).HandlerFunc(home).Name("home")

	return router

}
