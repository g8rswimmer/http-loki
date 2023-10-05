package replace

import (
	"fmt"

	"github.com/g8rswimmer/http-loki/internal/model"
	"github.com/google/uuid"
	"github.com/tidwall/sjson"
)

func UUID(_, resp string, bv model.BodyVariable) (string, error) {
	u := uuid.NewString()
	resp, err := sjson.Set(resp, bv.Path, u)
	if err != nil {
		return "", fmt.Errorf("response setting error %w", err)
	}
	return resp, nil
}
