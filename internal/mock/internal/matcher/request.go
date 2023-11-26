package matcher

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/g8rswimmer/http-loki/internal/httpx"
	"github.com/g8rswimmer/http-loki/internal/mock/internal/matcher/internal/body"
	"github.com/g8rswimmer/http-loki/internal/mock/internal/matcher/internal/query"
	"github.com/g8rswimmer/http-loki/internal/model"
)

func MockRequestMatch(req *httpx.Request, mockRequest model.Request) error {
	if err := matchQueryParameters(req, mockRequest); err != nil {
		fmt.Println(err.Error())
		return err
	}
	if err := matchBody(req, mockRequest); err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func matchQueryParameters(req *httpx.Request, mockRequest model.Request) error {
	switch {
	case len(req.QueryParameters) == 0 && len(mockRequest.QueryParameters) == 0:
		return nil
	case len(req.QueryParameters) != len(mockRequest.QueryParameters):
		return fmt.Errorf("request query parameters lenght does not match got: %d expected %d", len(req.QueryParameters), len(mockRequest.QueryParameters))
	default:
	}
	values, err := query.Validate(req.QueryParameters, mockRequest.QueryParameters)
	if err != nil {
		return fmt.Errorf("request query matching validation: %w", err)
	}
	for _, qp := range mockRequest.QueryParameters {
		v := values.Get(qp.Key)
		if len(v) == 0 || v != qp.Value {
			return fmt.Errorf("request query values do not match got: %s expected %s", v, qp.Value)
		}
	}
	return nil
}

func matchBody(req *httpx.Request, mockRequest model.Request) error {
	if req.Body == nil && mockRequest.Body == nil {
		return nil
	}
	rStr, err := body.Validate(req.EncodedBody, mockRequest.Validations)
	if err != nil {
		return fmt.Errorf("request body matching validation: %w", err)
	}
	var reqBody any
	if err := json.Unmarshal([]byte(rStr), &reqBody); err != nil {
		return fmt.Errorf("unable to marshal request body for comparing: %w", err)
	}
	if !reflect.DeepEqual(reqBody, mockRequest.Body) {
		return errors.New("request bodies to not match")
	}
	return nil
}
