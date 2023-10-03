package variable

import (
	"fmt"

	"github.com/g8rswimmer/http-loki/internal/variable/internal/replace"
)

type replacement func(string, string, string, []string) (string, error)

var replacements = map[string]replacement{
	"uuid": replace.UUID,
	"path": replace.Path,
}

func Replace(req, resp string, vars []Body) (string, error) {
	for _, v := range vars {
		repFunc, has := replacements[v.Func]
		if !has {
			return "", fmt.Errorf("variable validation func not found %s", v.Func)
		}
		var err error
		resp, err = repFunc(req, resp, v.Path, v.Args)
		if err != nil {
			return "", fmt.Errorf("variable validation %w", err)
		}
	}
	return resp, nil
}
