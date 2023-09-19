package mock

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"

	"github.com/g8rswimmer/http-loki/internal/model"
	"github.com/g8rswimmer/http-loki/internal/variable"
)

type pair struct {
	request   model.Request
	response  model.Response
	variables []variable.Body
}

type Handler struct {
	pairs []pair
}

func (h *Handler) Add(req model.Request, resp model.Response, v []variable.Body) {
	h.pairs = append(h.pairs, pair{request: req, response: resp, variables: v})
}

func (h *Handler) HTTPHandler(w http.ResponseWriter, r *http.Request) {
	p, err := h.requestPair(r)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(p.response.StatusCode)
	if p.response.Body == nil {
		w.Write([]byte("{}"))
		return
	}
	payload, _ := json.Marshal(p.response.Body)
	w.Write(payload)
}

func (h *Handler) requestPair(r *http.Request) (pair, error) {
	if len(h.pairs) == 0 {
		return pair{}, fmt.Errorf("no request pair")
	}
	var requestBody any
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	switch {
	case errors.Is(err, io.EOF):
	case err != nil:
		return pair{}, err
	default:
	}
	for _, p := range h.pairs {
		switch {
		case p.request.Body == nil && requestBody == nil:
			return p, nil
		case h.validateRequest(requestBody, p.request.Body, p):
			return p, nil
		default:
		}

	}
	return pair{}, fmt.Errorf("no request pair")
}

func (h *Handler) validateRequest(reqBody, mockBody any, p pair) bool {
	enc, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Println(err)
		return false
	}
	rStr, err := variable.Validate(string(enc), p.variables)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if err := json.Unmarshal([]byte(rStr), &reqBody); err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println("comparing")
	fmt.Println(rStr)
	return reflect.DeepEqual(reqBody, mockBody)
}
