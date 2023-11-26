package query

import (
	"github.com/g8rswimmer/http-loki/internal/mock/internal/matcher/internal/validate"
	"github.com/g8rswimmer/http-loki/internal/model"
)

func ignore(value string, qv model.QueryVariable) error {
	return validate.Ignore(value, qv.VariableParams)
}
