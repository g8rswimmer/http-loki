package matcher

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"

	"github.com/g8rswimmer/http-loki/internal/model"
	"github.com/g8rswimmer/http-loki/internal/variable"
)

type Request struct {
	body            any
	encodedBody     string
	queryParameters url.Values
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
		body:            requestBody,
		encodedBody:     string(enc),
		queryParameters: values,
	}, nil
}

func (r *Request) Match(req model.Request) (any, error) {
	if err := r.matchBody(req); err != nil {
		return nil, err
	}
	return r.body, nil
}

func (r *Request) matchBody(req model.Request) error {
	if r.body == nil && req.Body == nil {
		return nil
	}
	rStr, err := variable.Validate(r.encodedBody, req.Validations)
	if err != nil {
		return fmt.Errorf("request body matching validation: %w", err)
	}
	var reqBody any
	if err := json.Unmarshal([]byte(rStr), &reqBody); err != nil {
		return fmt.Errorf("unable to marshal request body for comparing: %w", err)
	}
	if !reflect.DeepEqual(reqBody, req.Body) {
		return errors.New("request bodies to not match")
	}
	return nil
}
