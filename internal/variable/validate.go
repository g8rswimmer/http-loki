package variable

import (
	"fmt"

	"github.com/tidwall/sjson"
)

type validation func(string, string, []string) error

var varMap = map[string]validation{
	"uuid":     validateUUID,
	"intRange": validateIntRange,
	"ignore":   validateIgnore,
	"regex":    validateRegEx,
}

func Validate(req string, vars []Body) (string, error) {
	for _, v := range vars {
		f, has := varMap[v.Func]
		if !has {
			return "", fmt.Errorf("variable validation func not found %s", v.Func)
		}
		err := f(req, v.Path, v.Args)
		if err != nil {
			return "", fmt.Errorf("variable validation %w", err)
		}
		req, err = sjson.Set(req, v.Path, "ignore")
		if err != nil {
			return "", fmt.Errorf("variable request setting %w", err)
		}
	}
	return req, nil
}
