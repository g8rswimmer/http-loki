package mock

import (
	"github.com/g8rswimmer/http-loki/internal/model"
)

type endpoints map[string]*Handler

func (e endpoints) add(m *model.Mock) bool {
	k := m.Method + ":" + m.Endpoint
	_, ok := e[k]
	if !ok {
		e[k] = &Handler{}
	}
	var reqVars []model.BodyVariable
	if m.Request.Body != nil {
		reqVars = m.Request.Validations
	}
	var respVars []model.BodyVariable
	if m.Response.Body != nil {
		respVars = m.Response.Replacements
	}
	e[k].Add(m.Request, reqVars, m.Response, respVars)
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
