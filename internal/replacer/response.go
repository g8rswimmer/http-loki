package replacer

import (
	"encoding/json"
	"fmt"

	"github.com/g8rswimmer/http-loki/internal/httpx"
	"github.com/g8rswimmer/http-loki/internal/model"
	"github.com/g8rswimmer/http-loki/internal/variable"
)

var emptyBody = []byte("{}")

func MockResponseReplace(req *httpx.Request, mockResponse model.Response) ([]byte, error) {
	if mockResponse.Body == nil {
		return emptyBody, nil
	}
	resp, err := json.Marshal(mockResponse.Body)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("response body marshal %w", err)
	}
	rStr, err := variable.Replace(req.EncodedBody, string(resp), mockResponse.Replacements)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("response body replace %w", err)
	}
	return []byte(rStr), nil
}
