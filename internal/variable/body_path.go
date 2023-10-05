package variable

import (
	"log"
	"strings"

	"github.com/g8rswimmer/http-loki/internal/model"
)

func BodyPaths(body any, currPath string, paths []model.BodyVariable) []model.BodyVariable {
	switch v := body.(type) {
	case map[string]any:
		paths = mapPaths(v, currPath, paths)
	default:
		log.Panicf("not supported %T %v", v, v)
	}
	return paths
}

func mapPaths(body map[string]any, currPath string, paths []model.BodyVariable) []model.BodyVariable {
	if len(currPath) > 0 {
		currPath += "."
	}
	for k, v := range body {
		switch value := v.(type) {
		case string:
			if strings.HasPrefix(value, "{{") && strings.HasSuffix(value, "}}") {
				vars := strings.TrimPrefix(value, "{{")
				vars = strings.TrimSuffix(vars, "}}")
				vars = strings.TrimSpace(vars)
				s := strings.Split(vars, ":")
				b := model.BodyVariable{
					Path: currPath + k,
					Func: s[0],
					Args: func() []string {
						if len(s) > 1 {
							return strings.Split(s[1], "|")
						}
						return []string{}
					}(),
				}
				paths = append(paths, b)
				body[k] = "ignore"
			}
		case map[string]any, []any:
			paths = BodyPaths(v, currPath+k, paths)
		default:
		}
	}
	return paths
}