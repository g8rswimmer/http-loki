package body

import (
	"fmt"
	"strings"

	"github.com/g8rswimmer/http-loki/internal/model"
	"github.com/tidwall/gjson"
)

func Ignore(req string, bv model.BodyVariable) error {
	result := gjson.Get(req, bv.Path)
	if result.Type != gjson.String {
		return nil
	}
	value := result.String()
	switch {
	case len(bv.Prefix) == 0:
	case !strings.HasPrefix(value, bv.Prefix):
		return fmt.Errorf("request does not have prefix %s %s", bv.Prefix, value)
	default:
	}
	switch {
	case len(bv.Suffix) == 0:
	case !strings.HasSuffix(value, bv.Suffix):
		return fmt.Errorf("request does not have suffix %s %s", bv.Suffix, value)
	default:
	}
	return nil
}
