package validate

import "github.com/g8rswimmer/http-loki/internal/model"

func Ignore(_ string, _ model.BodyVariable) error {
	return nil
}
