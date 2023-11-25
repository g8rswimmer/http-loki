package query

import (
	"net/url"

	"github.com/g8rswimmer/http-loki/internal/model"
)

func Validate(values url.Values, vars []model.QueryVariable) (url.Values, error) {
	return values, nil
}
