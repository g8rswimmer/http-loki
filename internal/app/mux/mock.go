package mux

import (
	"github.com/g8rswimmer/http-loki/internal/model"
	"github.com/gorilla/mux"
)

func add(m *model.Mock, ep endpoints, route *mux.Router) {
	if has := ep.add(m); !has {
		route.Methods(m.Method).Path(m.Endpoint).HandlerFunc(ep.handler(m).HTTPHandler)
	}
}
