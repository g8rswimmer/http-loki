package body

import (
	"github.com/g8rswimmer/http-loki/internal/matcher/internal/validate"
	"github.com/g8rswimmer/http-loki/internal/model"
	"github.com/tidwall/gjson"
)

func regex(req string, bv model.BodyVariable) error {
	result := gjson.Get(req, bv.Path)
	return validate.RegEx(result.Str, bv.VariableParams)
}
