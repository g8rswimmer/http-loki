package query

import (
	"fmt"
	"net/url"

	"github.com/g8rswimmer/http-loki/internal/model"
)

const validationValue = "{{ validation }}"

type validation func(string, model.QueryVariable) error

var validations = map[string]validation{
	"uuid":     uuid,
	"ignore":   ignore,
	"intRange": intRange,
	"regex":    regex,
	"oneOf":    oneOf,
}

func Validate(values url.Values, params []model.QueryParameter) (url.Values, error) {
	varValues := url.Values{}
	for k := range values {
		varValues.Set(k, values.Get(k))
	}
	for _, param := range params {
		val := param.Validation
		if val == nil {
			continue
		}
		valFunc, has := validations[val.Func]
		if !has {
			return varValues, fmt.Errorf("variable validation func not found %s", val.Func)
		}
		err := valFunc(values.Get(param.Key), *val)
		if err != nil {
			return varValues, fmt.Errorf("variable validation %w", err)
		}
		varValues.Set(param.Key, validationValue)
	}
	fmt.Println(varValues)
	return varValues, nil
}
