package body

import (
	"fmt"

	"github.com/g8rswimmer/http-loki/internal/mock/internal/matcher/internal/validate"
	"github.com/g8rswimmer/http-loki/internal/model"
	"github.com/tidwall/gjson"
)

func IntRange(req string, bv model.BodyVariable) error {
	result := gjson.Get(req, bv.Path)
	switch {
	case result.Type == gjson.Number:
		return validate.IntRange(int(result.Int()), bv.VariableParams)
	case result.IsArray():
		for _, r := range result.Array() {
			if r.Type != gjson.Number {
				return fmt.Errorf("request path %s is not a number %v", bv.Path, r.String())
			}
			if err := validate.IntRange(int(r.Int()), bv.VariableParams); err != nil {
				return err
			}
		}
		return nil
	default:
		return fmt.Errorf("request path %s is not a number %v", bv.Path, result.String())
	}
}
