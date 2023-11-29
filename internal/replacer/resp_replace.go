package replacer

import (
	"fmt"

	"github.com/g8rswimmer/http-loki/internal/httpx"
	"github.com/g8rswimmer/http-loki/internal/model"
	"github.com/g8rswimmer/http-loki/internal/replacer/internal/replace"
)

type replacement func(*httpx.Request, string, model.BodyVariable) (string, error)

var replacements = map[string]replacement{
	"uuid":     replace.UUID,
	"path":     replace.Path,
	"currTime": replace.CurrentTime,
}

func replaceVars(req *httpx.Request, resp string, vars []model.BodyVariable) (string, error) {
	for _, v := range vars {
		repFunc, has := replacements[v.Func]
		if !has {
			return "", fmt.Errorf("variable validation func not found %s", v.Func)
		}
		var err error
		resp, err = repFunc(req, resp, v)
		if err != nil {
			return "", fmt.Errorf("variable validation %w", err)
		}
	}
	return resp, nil
}
