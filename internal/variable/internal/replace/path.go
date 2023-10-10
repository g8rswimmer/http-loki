package replace

import (
	"fmt"

	"github.com/g8rswimmer/http-loki/internal/model"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func Path(req, resp string, bv model.BodyVariable) (string, error) {
	if len(bv.Args) != 1 {
		return "", fmt.Errorf("response args not length 1 %d", len(bv.Args))
	}
	result := gjson.Get(req, bv.Args[0])
	var value any
	switch {
	case result.Type == gjson.String:
		value = bv.Prefix + result.Str + bv.Suffix
	default:
		value = result.Value()
	}
	resp, err := sjson.Set(resp, bv.Path, value)
	if err != nil {
		return "", fmt.Errorf("response setting error %w", err)
	}
	return resp, nil
}
