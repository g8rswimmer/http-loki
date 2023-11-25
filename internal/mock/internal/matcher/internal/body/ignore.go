package body

import (
	"github.com/g8rswimmer/http-loki/internal/mock/internal/matcher/internal/validate"
	"github.com/g8rswimmer/http-loki/internal/model"
	"github.com/tidwall/gjson"
)

func ignore(req string, bv model.BodyVariable) error {
	result := gjson.Get(req, bv.Path)
	if result.Type != gjson.String {
		return nil
	}
	return validate.Ignore(result.Str, bv.VariableParams)
}
