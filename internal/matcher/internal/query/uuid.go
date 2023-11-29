package query

import (
	"github.com/g8rswimmer/http-loki/internal/matcher/internal/validate"
	"github.com/g8rswimmer/http-loki/internal/model"
)

func uuid(value string, qv model.QueryVariable) error {
	return validate.UUID(value, qv.VariableParams)
}
