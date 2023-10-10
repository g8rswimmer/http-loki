package validate

import (
	"fmt"
	"strings"

	"github.com/g8rswimmer/http-loki/internal/model"
	"github.com/go-playground/validator/v10"
	"github.com/tidwall/gjson"
)

func UUID(req string, bv model.BodyVariable) error {
	result := gjson.Get(req, bv.Path)
	switch {
	case result.Type == gjson.String:
		return validateUUID(result.Str, bv.Prefix, bv.Suffix)
	case result.IsArray():
		for _, r := range result.Array() {
			if r.Type != gjson.String {
				return fmt.Errorf("request value is not a string %v", r.Raw)
			}
			if err := validateUUID(r.Str, bv.Prefix, bv.Suffix); err != nil {
				return err
			}
		}
		return nil
	default:
		return fmt.Errorf("request value is not a string %v", result.Raw)
	}
}

func validateUUID(value, prefix, suffix string) error {
	switch {
	case len(prefix) == 0:
	case !strings.HasPrefix(value, prefix):
		return fmt.Errorf("request does not have prefix %s %s", prefix, value)
	default:
		value = strings.TrimPrefix(value, prefix)
	}
	switch {
	case len(suffix) == 0:
	case !strings.HasSuffix(value, suffix):
		return fmt.Errorf("request does not have prefix %s %s", suffix, value)
	default:
		value = strings.TrimSuffix(value, suffix)
	}
	return validator.New().Var(value, "uuid4")
}
