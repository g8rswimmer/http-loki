package mock

import (
	"fmt"

	"github.com/g8rswimmer/http-loki/internal/model"
	"github.com/g8rswimmer/http-loki/internal/variable"
)

type endpoints map[string]*Handler

func (e endpoints) add(m *model.Mock) bool {
	k := m.Method + ":" + m.Endpoint
	_, ok := e[k]
	if !ok {
		e[k] = &Handler{}
	}
	var vars []variable.Body
	if m.Request.Body != nil {
		vars = variable.BodyPaths(m.Request.Body, "", []variable.Body{})
	}
	fmt.Printf("%+v\n", vars)
	e[k].Add(m.Request, m.Response, vars)
	return ok
}

func (e endpoints) handler(m *model.Mock) *Handler {
	k := m.Method + ":" + m.Endpoint
	h, ok := e[k]
	if !ok {
		return &Handler{}
	}
	return h
}
