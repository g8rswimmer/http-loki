package variable

import (
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/tidwall/gjson"
)

func validateIntRange(req string, path string, args []string) error {
	result := gjson.Get(req, path)
	if result.Type != gjson.Number {
		return fmt.Errorf("request path %s is not a number", path)
	}
	if len(args) != 2 {
		return fmt.Errorf("request arg length is not two %d", len(args))
	}
	low, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("request arg is not a number %s", args[0])
	}
	high, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("request arg is not a number %s", args[1])
	}
	if low > high {
		low, high = high, low
	}
	return validator.New().Var(result.Int(), fmt.Sprintf("gte=%d,lte=%d", low, high))
}