package request

import "reflect"

func Compare(reqBody, mockBody any) bool {
	return reflect.DeepEqual(reqBody, mockBody)
}
