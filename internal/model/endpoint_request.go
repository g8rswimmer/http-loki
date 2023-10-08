package model

type Request struct {
	Body        any            `json:"body"`
	Validations []BodyVariable `json:"body_validations"`
}

type Response struct {
	StatusCode   int            `json:"status_code"`
	Body         any            `json:"body"`
	Replacements []BodyVariable `json:"body_replacements"`
}

type Mock struct {
	Method   string   `json:"method"`
	Endpoint string   `json:"endpoint"`
	Request  Request  `json:"request"`
	Response Response `json:"response"`
}
