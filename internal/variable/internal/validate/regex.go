package validate

import (
	"fmt"
	"regexp"

	"github.com/g8rswimmer/http-loki/internal/model"
	"github.com/tidwall/gjson"
)

func RegEx(req string, bv model.BodyVariable) error {
	if len(bv.Args) != 1 {
		return fmt.Errorf("request arg length is not two %d", len(bv.Args))
	}
	result := gjson.Get(req, bv.Path)
	match, err := regexp.MatchString(bv.Args[0], result.String())
	switch {
	case err != nil:
		return fmt.Errorf("request reg exp error %w", err)
	case !match:
		return fmt.Errorf("request reg exp not a match %s %s", bv.Args[0], result.String())
	default:
	}
	return nil
}
