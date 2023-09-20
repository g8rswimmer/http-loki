package replace

import (
	"fmt"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func Path(req, resp string, path string, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("response args not length 1 %d", len(args))
	}
	result := gjson.Get(req, args[0])
	resp, err := sjson.Set(resp, path, result.Value())
	if err != nil {
		return "", fmt.Errorf("response setting error %w", err)
	}
	return resp, nil
}
