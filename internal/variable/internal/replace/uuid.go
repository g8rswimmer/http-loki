package replace

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/tidwall/sjson"
)

func UUID(_, resp string, path string, _ []string) (string, error) {
	u := uuid.NewString()
	resp, err := sjson.Set(resp, path, u)
	if err != nil {
		return "", fmt.Errorf("response setting error %w", err)
	}
	return resp, nil
}
