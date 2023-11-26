package body

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/g8rswimmer/http-loki/internal/httpx"
	"github.com/g8rswimmer/http-loki/internal/model"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

const (
	valueArg = 0
	keyArg   = 1
	typeArg  = 2
	maxArgs  = 3

	stringType = "string"
	intType    = "int"
	floatType  = "float"
	boolType   = "bool"

	reqBodyValue  = "body"
	reqQueryValue = "query"
)

func Path(req *httpx.Request, resp string, bv model.BodyVariable) (string, error) {
	if len(bv.Args) < 2 {
		return "", fmt.Errorf("response args not length at least 2 %d", len(bv.Args))
	}

	var value any
	var err error
	switch bv.Args[valueArg] {
	case reqBodyValue:
		value = bodyValue(req.EncodedBody, bv.VariableParams)
	case reqQueryValue:
		value, err = queryValue(req.QueryParameters, bv.VariableParams)
	default:
		err = fmt.Errorf("response value arg not supported %s", bv.Args[valueArg])
	}
	if err != nil {
		return "", err
	}

	resp, err = sjson.Set(resp, bv.Path, value)
	if err != nil {
		return "", fmt.Errorf("response setting error %w", err)
	}
	return resp, nil
}

func bodyValue(req string, params model.VariableParams) any {
	result := gjson.Get(req, params.Args[keyArg])
	var value any
	switch {
	case result.Type == gjson.String:
		value = params.Prefix + result.Str + params.Suffix
	default:
		value = result.Value()
	}
	return value
}

func queryValue(values url.Values, params model.VariableParams) (any, error) {
	if !values.Has(params.Args[keyArg]) {
		return nil, fmt.Errorf("request query key not present %s", params.Args[keyArg])
	}
	t := stringType
	if len(params.Args) == maxArgs {
		t = params.Args[typeArg]
	}

	v := values.Get(params.Args[keyArg])
	var value any
	switch t {
	case intType:
		i, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("request query value not a number %s: %w", v, err)
		}
		value = i
	case floatType:
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, fmt.Errorf("request query value not a number %s: %w", v, err)
		}
		value = f
	case boolType:
		b, err := strconv.ParseBool(v)
		if err != nil {
			return nil, fmt.Errorf("request query value not a bool %s: %w", v, err)
		}
		value = b
	default:
		value = params.Prefix + v + params.Suffix
	}

	return value, nil
}
