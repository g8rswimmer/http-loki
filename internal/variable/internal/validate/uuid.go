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
	value := result.String()
	switch {
	case len(bv.Prefix) == 0:
	case !strings.HasPrefix(value, bv.Prefix):
		return fmt.Errorf("request does not have prefix %s %s", bv.Prefix, value)
	default:
		value = strings.TrimPrefix(value, bv.Prefix)
	}
	switch {
	case len(bv.Suffix) == 0:
	case !strings.HasSuffix(value, bv.Suffix):
		return fmt.Errorf("request does not have prefix %s %s", bv.Suffix, value)
	default:
		value = strings.TrimSuffix(value, bv.Suffix)
	}
	return validator.New().Var(value, "uuid4")
}
