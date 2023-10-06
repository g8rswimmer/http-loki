package variable

import (
	"log"

	"github.com/g8rswimmer/http-loki/internal/model"
)

func BodyPaths(body any, currPath string, bvs []model.BodyVariable) []model.BodyVariable {
	switch v := body.(type) {
	case map[string]any:
		bvs = mapPaths(v, currPath, bvs)
	default:
		log.Panicf("not supported %T %v", v, v)
	}
	return bvs
}

func mapPaths(body map[string]any, currPath string, bvs []model.BodyVariable) []model.BodyVariable {
	if len(currPath) > 0 {
		currPath += "."
	}
	for k, v := range body {
		switch value := v.(type) {
		case string:
			if bv, has := model.BodyVariableFromString(currPath+k, value); has {
				bvs = append(bvs, bv)
				body[k] = "ignore"
			}
		case map[string]any, []any:
			bvs = BodyPaths(v, currPath+k, bvs)
		default:
		}
	}
	return bvs
}
