package replace

import (
	"fmt"
	"time"

	"github.com/g8rswimmer/http-loki/internal/model"
	"github.com/tidwall/sjson"
)

func CurrentTime(_, resp string, bv model.BodyVariable) (string, error) {
	if len(bv.Args) != 1 {
		return "", fmt.Errorf("response args not length 1 %d", len(bv.Args))
	}
	var layout string
	switch bv.Args[0] {
	case "RFC3339":
		layout = time.RFC3339
	default:
		layout = bv.Args[0]
	}
	t := time.Now().Format(layout)
	value := bv.Prefix + t + bv.Suffix
	resp, err := sjson.Set(resp, bv.Path, value)
	if err != nil {
		return "", fmt.Errorf("response setting error %w", err)
	}
	return resp, nil
}
