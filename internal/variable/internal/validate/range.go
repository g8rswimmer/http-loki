package validate

import (
	"fmt"
	"strconv"

	"github.com/g8rswimmer/http-loki/internal/model"
	"github.com/go-playground/validator/v10"
	"github.com/tidwall/gjson"
)

func IntRange(req string, bv model.BodyVariable) error {
	if len(bv.Args) != 2 {
		return fmt.Errorf("request arg length is not two %d", len(bv.Args))
	}
	result := gjson.Get(req, bv.Path)
	if result.Type != gjson.Number {
		return fmt.Errorf("request path %s is not a number", bv.Path)
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
	return validator.New().Var(result.Int(), fmt.Sprintf("gte=%d,lte=%d", low, high))
}
