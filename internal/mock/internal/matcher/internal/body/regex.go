package body

import (
	"github.com/g8rswimmer/http-loki/internal/mock/internal/matcher/internal/validate"
	"github.com/g8rswimmer/http-loki/internal/model"
	"github.com/tidwall/gjson"
)

func RegEx(req string, bv model.BodyVariable) error {
	result := gjson.Get(req, bv.Path)
	return validate.RegEx(result.Str, bv.VariableParams)
}
