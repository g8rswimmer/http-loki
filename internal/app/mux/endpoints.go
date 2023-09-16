package mux

import (
	"github.com/g8rswimmer/http-loki/internal/mock"
	"github.com/g8rswimmer/http-loki/internal/model"
)

type endpoints map[string]*mock.Handler

func (e endpoints) add(m *model.Mock) bool {
	k := m.Method + ":" + m.Endpoint
	_, ok := e[k]
	if !ok {
		e[k] = &mock.Handler{}
	}
	e[k].Add(m.Request, m.Response)
	return ok
}

func (e endpoints) handler(m *model.Mock) *mock.Handler {
	k := m.Method + ":" + m.Endpoint
	h, ok := e[k]
	if !ok {
		return &mock.Handler{}
	}
	return h
}
