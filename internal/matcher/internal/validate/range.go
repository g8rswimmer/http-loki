package validate

import (
	"fmt"
	"strconv"

	"github.com/g8rswimmer/http-loki/internal/model"
)

func IntRange(value int, params model.VariableParams) error {
	if len(params.Args) != 2 {
		return fmt.Errorf("request arg length is not two %d", len(params.Args))
	}
	low, err := strconv.Atoi(params.Args[0])
	if err != nil {
		return fmt.Errorf("request arg is not a number %s", params.Args[0])
	}
	high, err := strconv.Atoi(params.Args[1])
	if err != nil {
		return fmt.Errorf("request arg is not a number %s", params.Args[1])
	}
	if low > high {
		low, high = high, low
	}
	if value < low || value > high {
		return fmt.Errorf("request arg %d is not between %d and %d", value, low, high)
	}
	return nil
}
