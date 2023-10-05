package validate

import (
	"github.com/g8rswimmer/http-loki/internal/model"
	"github.com/go-playground/validator/v10"
	"github.com/tidwall/gjson"
)

func UUID(req string, bv model.BodyVariable) error {
	result := gjson.Get(req, bv.Path)
	return validator.New().Var(result.String(), "uuid4")
}
