package validate

import (
	"fmt"
	"strings"

	"github.com/g8rswimmer/http-loki/internal/model"
)

func Ignore(value string, params model.VariableParams) error {
	switch {
	case len(params.Prefix) == 0:
	case !strings.HasPrefix(value, params.Prefix):
		return fmt.Errorf("request does not have prefix %s %s", params.Prefix, value)
	default:
	}
	switch {
	case len(params.Suffix) == 0:
	case !strings.HasSuffix(value, params.Suffix):
		return fmt.Errorf("request does not have suffix %s %s", params.Suffix, value)
	default:
	}
	return nil
}
