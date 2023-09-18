package variable

import (
	"log"
	"strings"
)

type Body struct {
	Path     string
	Variable string
}

func BodyPaths(body any, currPath string, paths []Body) []Body {
	switch v := body.(type) {
	case map[string]any:
		paths = mapPaths(v, currPath, paths)
	default:
		log.Panicf("not supported %T %v", v, v)
	}
	return paths
}

func mapPaths(body map[string]any, currPath string, paths []Body) []Body {
	if len(currPath) > 0 {
		currPath += "."
	}
	for k, v := range body {
		switch value := v.(type) {
		case string:
			if strings.HasPrefix(value, "{{") && strings.HasSuffix(value, "}}") {
				paths = append(paths, Body{Path: currPath + k, Variable: value})
				delete(body, k)
			}
		case map[string]any, []any:
			paths = BodyPaths(v, currPath+k, paths)
		default:
		}
	}
	return paths
}
