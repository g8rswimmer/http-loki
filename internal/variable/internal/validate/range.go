package validate

import (
	"fmt"
	"strconv"

	"github.com/g8rswimmer/http-loki/internal/model"
	"github.com/tidwall/gjson"
)

func IntRange(req string, bv model.BodyVariable) error {
	if len(bv.Args) != 2 {
		return fmt.Errorf("request arg length is not two %d", len(bv.Args))
	}
	low, err := strconv.Atoi(bv.Args[0])
	if err != nil {
		return fmt.Errorf("request arg is not a number %s", bv.Args[0])
	}
	high, err := strconv.Atoi(bv.Args[1])
	if err != nil {
		return fmt.Errorf("request arg is not a number %s", bv.Args[1])
	}
	if low > high {
		low, high = high, low
	}
	result := gjson.Get(req, bv.Path)
	switch {
	case result.Type == gjson.Number:
		return validateIntRange(int(result.Int()), low, high)
	case result.IsArray():
		for _, r := range result.Array() {
			if r.Type != gjson.Number {
				return fmt.Errorf("request path %s is not a number %v", bv.Path, r.String())
			}
			if err := validateIntRange(int(r.Int()), low, high); err != nil {
				return err
			}
		}
		return nil
	default:
		return fmt.Errorf("request path %s is not a number %v", bv.Path, result.String())
	}
}

func validateIntRange(value, low, high int) error {
	if value < low || value > high {
		return fmt.Errorf("request arg %d is not between %d and %d", value, low, high)
	}
	return nil
}
