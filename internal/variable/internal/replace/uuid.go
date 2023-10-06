package replace

import (
	"fmt"

	"github.com/g8rswimmer/http-loki/internal/model"
	"github.com/google/uuid"
	"github.com/tidwall/sjson"
)

func UUID(_, resp string, bv model.BodyVariable) (string, error) {
	u := uuid.NewString()
	value := bv.Prefix + u + bv.Suffix
	resp, err := sjson.Set(resp, bv.Path, value)
	if err != nil {
		return "", fmt.Errorf("response setting error %w", err)
	}
	return resp, nil
}
