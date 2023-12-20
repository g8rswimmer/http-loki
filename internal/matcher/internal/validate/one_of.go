package validate

import (
	"fmt"
	"strings"

	"github.com/g8rswimmer/http-loki/internal/model"
)

func OneOf(value string, params model.VariableParams) error {
	for _, arg := range params.Args {
		if value == arg {
			return nil
		}
	}
	return fmt.Errorf("value (%s) was not one of the supplied args: %v", value, strings.Join(params.Args, ", "))
}
