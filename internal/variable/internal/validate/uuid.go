package validate

import (
	"github.com/go-playground/validator/v10"
	"github.com/tidwall/gjson"
)

func UUID(req string, path string, _ []string) error {
	result := gjson.Get(req, path)
	return validator.New().Var(result.String(), "uuid4")
}
