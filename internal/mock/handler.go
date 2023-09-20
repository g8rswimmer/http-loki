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
	request      model.Request
	response     model.Response
	requestVars  []variable.Body
	responseVars []variable.Body
}

type Handler struct {
	pairs []pair
}

func (h *Handler) Add(req model.Request, reqVars []variable.Body, resp model.Response, respVars []variable.Body) {
	p := pair{
		request:      req,
		response:     resp,
		requestVars:  reqVars,
		responseVars: respVars,
	}
	h.pairs = append(h.pairs, p)
}

func (h *Handler) HTTPHandler(w http.ResponseWriter, r *http.Request) {
	rb, p, err := h.requestPair(r)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Add("content-type", "application/json")
	statusCode := p.response.StatusCode
	respBody := "{}"
	if p.response.Body != nil {
		r, err := h.replaceResponse(rb, p.response.Body, p)
		if err != nil {
			statusCode = http.StatusMethodNotAllowed
		}
		respBody = r
	}

	w.WriteHeader(statusCode)
	w.Write([]byte(respBody))
}

func (h *Handler) requestPair(r *http.Request) (any, pair, error) {
	if len(h.pairs) == 0 {
		return nil, pair{}, fmt.Errorf("no request pair")
	}
	var requestBody any
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	switch {
	case errors.Is(err, io.EOF):
	case err != nil:
		return nil, pair{}, err
	default:
	}
	for _, p := range h.pairs {
		switch {
		case p.request.Body == nil && requestBody == nil:
			return requestBody, p, nil
		case h.validateRequest(requestBody, p.request.Body, p):
			return requestBody, p, nil
		default:
		}

	}
	return nil, pair{}, fmt.Errorf("no request pair")
}

func (h *Handler) validateRequest(reqBody, mockBody any, p pair) bool {
	enc, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Println(err)
		return false
	}
	rStr, err := variable.Validate(string(enc), p.requestVars)
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

func (h *Handler) replaceResponse(requestBody, responseBody any, p pair) (string, error) {
	resp, err := json.Marshal(responseBody)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("response body marshal %w", err)
	}
	req, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("request body marshal %w", err)
	}
	rStr, err := variable.Replace(string(req), string(resp), p.responseVars)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("response body replace %w", err)
	}
	return rStr, nil
}
