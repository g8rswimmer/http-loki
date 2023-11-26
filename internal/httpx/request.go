package httpx

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Request struct {
	Body            any
	EncodedBody     string
	QueryParameters url.Values
}

func NewRequest(req *http.Request) (*Request, error) {

	var requestBody any
	err := json.NewDecoder(req.Body).Decode(&requestBody)
	switch {
	case errors.Is(err, io.EOF):
	case err != nil:
		return nil, fmt.Errorf("unable to decode request body [%s:%s]: %w", req.Method, req.URL.Path, err)
	default:
	}

	enc, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("unable to encode request body [%s:%s]: %w", req.Method, req.URL.Path, err)
	}

	values, err := url.ParseQuery(req.URL.RawQuery)
	if err != nil {
		return nil, fmt.Errorf("unable to parse query [%s:%s]: %w", req.Method, req.URL.Path, err)
	}
	return &Request{
		Body:            requestBody,
		EncodedBody:     string(enc),
		QueryParameters: values,
	}, nil
}
