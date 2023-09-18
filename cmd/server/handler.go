package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	router := mux.NewRouter()

	router.Methods(http.MethodGet).Path("/").HandlerFunc(home).Name("home")

	return router

}
