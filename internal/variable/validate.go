package variable

import (
	"fmt"

	"github.com/g8rswimmer/http-loki/internal/model"
	"github.com/g8rswimmer/http-loki/internal/variable/internal/validate"
	"github.com/tidwall/sjson"
)

const validationValue = "{{ validation }}"

type validation func(string, model.BodyVariable) error

var validations = map[string]validation{
	"uuid":     validate.UUID,
	"intRange": validate.IntRange,
	"ignore":   validate.Ignore,
	"regex":    validate.RegEx,
}

func Validate(req string, vars []model.BodyVariable) (string, error) {
	for _, v := range vars {
		valFunc, has := validations[v.Func]
		if !has {
			return "", fmt.Errorf("variable validation func not found %s", v.Func)
		}
		err := valFunc(req, v)
		if err != nil {
			return "", fmt.Errorf("variable validation %w", err)
		}
		req, err = sjson.Set(req, v.Path, validationValue)
		if err != nil {
			return "", fmt.Errorf("variable request setting %w", err)
		}
	}
	return req, nil
}
