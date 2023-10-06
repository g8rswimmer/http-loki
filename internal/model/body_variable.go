package model

import "strings"

const (
	bodyVariableFunc  = 0
	bodyVariableArgs  = 1
	argsSplit         = "|"
	funcSplit         = ":"
	bodyVariableStart = "{{"
	bodyVariableEnd   = "}}"
)

type BodyVariable struct {
	Path   string
	Func   string
	Args   []string
	Prefix string
	Suffix string
}

func BodyVariableFromString(path, field string) (BodyVariable, bool) {
	if !strings.Contains(field, bodyVariableStart) || !strings.Contains(field, bodyVariableEnd) {
		return BodyVariable{}, false
	}
	bv := BodyVariable{}

	p := strings.Split(field, bodyVariableStart)
	switch {
	case len(p) > 1:
		bv.Prefix = p[0]
		field = p[1]
	default:
		field = p[0]
	}

	s := strings.Split(field, bodyVariableEnd)
	field = s[0]
	if len(s) > 1 {
		bv.Suffix = s[1]
	}

	field = strings.TrimSpace(field)
	vars := strings.Split(field, funcSplit)
	bv.Path = path
	bv.Func = vars[bodyVariableFunc]
	bv.Args = func() []string {
		if len(vars) > bodyVariableArgs {
			return strings.Split(vars[bodyVariableArgs], argsSplit)
		}
		return []string{}
	}()

	return bv, true
}
