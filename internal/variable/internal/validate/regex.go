package validate

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/g8rswimmer/http-loki/internal/model"
	"github.com/tidwall/gjson"
)

func RegEx(req string, bv model.BodyVariable) error {
	if len(bv.Args) != 1 {
		return fmt.Errorf("request arg length is not two %d", len(bv.Args))
	}
	result := gjson.Get(req, bv.Path)
	value := result.String()
	switch {
	case len(bv.Prefix) == 0:
	case !strings.HasPrefix(value, bv.Prefix):
		return fmt.Errorf("request does not have prefix %s %s", bv.Prefix, value)
	default:
		value = strings.TrimPrefix(value, bv.Prefix)
	}
	switch {
	case len(bv.Suffix) == 0:
	case !strings.HasSuffix(value, bv.Suffix):
		return fmt.Errorf("request does not have prefix %s %s", bv.Suffix, value)
	default:
		value = strings.TrimSuffix(value, bv.Suffix)
	}
	match, err := regexp.MatchString(bv.Args[0], value)
	switch {
	case err != nil:
		return fmt.Errorf("request reg exp error %w", err)
	case !match:
		return fmt.Errorf("request reg exp not a match %s %s", bv.Args[0], value)
	default:
	}
	return nil
}
