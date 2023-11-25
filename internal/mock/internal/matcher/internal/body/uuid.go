package body

import (
	"fmt"

	"github.com/g8rswimmer/http-loki/internal/mock/internal/matcher/internal/validate"
	"github.com/g8rswimmer/http-loki/internal/model"
	"github.com/tidwall/gjson"
)

func uuid(req string, bv model.BodyVariable) error {
	result := gjson.Get(req, bv.Path)
	switch {
	case result.Type == gjson.String:
		return validate.UUID(result.Str, bv.VariableParams)
	case result.IsArray():
		for _, r := range result.Array() {
			if r.Type != gjson.String {
				return fmt.Errorf("request value is not a string %v", r.Raw)
			}
			if err := validate.UUID(r.Str, bv.VariableParams); err != nil {
				return err
			}
		}
		return nil
	default:
		return fmt.Errorf("request value is not a string %v", result.Raw)
	}
}
