package query

import (
	"fmt"
	"strconv"

	"github.com/g8rswimmer/http-loki/internal/mock/internal/matcher/internal/validate"
	"github.com/g8rswimmer/http-loki/internal/model"
)

func intRange(value string, qv model.QueryVariable) error {
	vint, err := strconv.Atoi(value)
	if err != nil {
		return fmt.Errorf("query arg is not a number %s", value)
	}
	return validate.IntRange(vint, qv.VariableParams)
}
