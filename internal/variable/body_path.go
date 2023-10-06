package variable

import (
	"log"

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
			if bv, has := model.BodyVariableFromString(currPath+k, value); has {
				paths = append(paths, bv)
				body[k] = "ignore"
			}
		case map[string]any, []any:
			paths = BodyPaths(v, currPath+k, paths)
		default:
		}
	}
	return paths
}
