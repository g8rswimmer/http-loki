package model

type BodyVariable struct {
	Path   string   `json:"json_path"`
	Func   string   `json:"func"`
	Args   []string `json:"args"`
	Prefix string   `json:"prefix"`
	Suffix string   `json:"suffix"`
}

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
