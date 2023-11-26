package validate

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/g8rswimmer/http-loki/internal/model"
)

func RegEx(value string, params model.VariableParams) error {
	if len(params.Args) != 1 {
		return fmt.Errorf("request arg length is not two %d", len(params.Args))
	}
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
	match, err := regexp.MatchString(params.Args[0], value)
	switch {
	case err != nil:
		return fmt.Errorf("request reg exp error %w", err)
	case !match:
		return fmt.Errorf("request reg exp not a match %s %s", params.Args[0], value)
	default:
	}
	return nil
}
