package variable

import (
	"fmt"
	"regexp"

	"github.com/tidwall/gjson"
)

func validateRegEx(req string, path string, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("request arg length is not two %d", len(args))
	}
	result := gjson.Get(req, path)
	match, err := regexp.MatchString(args[0], result.String())
	switch {
	case err != nil:
		return fmt.Errorf("request reg exp error %w", err)
	case !match:
		return fmt.Errorf("request reg exp not a match %s %s", args[0], result.String())
	default:
	}
	return nil
}
