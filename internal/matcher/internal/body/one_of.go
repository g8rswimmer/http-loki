package body

import (
	"github.com/g8rswimmer/http-loki/internal/matcher/internal/validate"
	"github.com/g8rswimmer/http-loki/internal/model"
	"github.com/tidwall/gjson"
)

func oneOf(value string, bv model.BodyVariable) error {
	realValue := gjson.Get(value, bv.Path)
	return validate.OneOf(realValue.String(), bv.VariableParams)
}
