package validate

import (
	"fmt"
	"strings"

	"github.com/g8rswimmer/http-loki/internal/model"
	"github.com/go-playground/validator/v10"
)

func UUID(value string, params model.VariableParams) error {
	switch {
	case len(params.Prefix) == 0:
	case !strings.HasPrefix(value, params.Prefix):
		return fmt.Errorf("request does not have prefix %s %s", params.Prefix, value)
	default:
		value = strings.TrimPrefix(value, params.Prefix)
	}
	switch {
	case len(params.Suffix) == 0:
	case !strings.HasSuffix(value, params.Suffix):
		return fmt.Errorf("request does not have prefix %s %s", params.Suffix, value)
	default:
		value = strings.TrimSuffix(value, params.Suffix)
	}
	return validator.New().Var(value, "uuid4")
}
