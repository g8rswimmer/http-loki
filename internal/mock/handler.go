package mock

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/g8rswimmer/http-loki/internal/httpx"
	"github.com/g8rswimmer/http-loki/internal/matcher"
	"github.com/g8rswimmer/http-loki/internal/model"
	"github.com/g8rswimmer/http-loki/internal/replacer"
)

const (
	errorStatusCode = 477
)

type pair struct {
	request  model.Request
	response model.Response
}

type Handler struct {
	pairs          []pair
	defaultRespone *model.Response
}

func (h *Handler) Add(req model.Request, resp model.Response) {
	p := pair{
		request:  req,
		response: resp,
	}
	h.pairs = append(h.pairs, p)
}

func (h *Handler) AddDefaultResponse(resp model.Response) {
	h.defaultRespone = &resp
}

func (h *Handler) HTTPHandler(w http.ResponseWriter, r *http.Request) {
	request, err := httpx.NewRequest(r)
	if err != nil {
		fmt.Println(err)
		h.errorResponse(w, errorStatusCode, "mock request error", err)
		return
	}
	resp, err := h.findResponse(request)
	if err != nil {
		fmt.Println(err)
		h.errorResponse(w, errorStatusCode, "mock request error", err)
		return
	}

	respBody, err := replacer.MockResponseReplace(request, *resp)
	if err != nil {
		h.errorResponse(w, errorStatusCode, "mock response error", err)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(respBody)
}

func (h Handler) errorResponse(w http.ResponseWriter, statusCode int, msg string, err error) {
	body := struct {
		Msg string `json:"msg"`
		Err string `json:"error,omitempty"`
	}{
		Msg: msg,
		Err: func() string {
			if err != nil {
				return err.Error()
			}
			return ""
		}(),
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(body)
}

func (h *Handler) findResponse(request *httpx.Request) (*model.Response, error) {
	resp := h.defaultRespone
	for _, p := range h.pairs {
		if err := matcher.MockRequestMatch(request, p.request); err == nil {
			resp = &p.response
			break
		}
	}
	switch {
	case resp == nil:
		return nil, fmt.Errorf("no request pair")
	default:
		return resp, nil
	}

}
