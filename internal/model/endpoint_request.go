package model

type BodyVariable struct {
	Path   string   `json:"json_path"`
	Func   string   `json:"func"`
	Args   []string `json:"args"`
	Prefix string   `json:"prefix"`
	Suffix string   `json:"suffix"`
}

type QueryVariable struct {
	Func   string   `json:"func"`
	Args   []string `json:"args"`
	Prefix string   `json:"prefix"`
	Suffix string   `json:"suffix"`
}

type QueryParameter struct {
	Key        string        `json:"key"`
	Value      string        `json:"value"`
	Validation QueryVariable `json:"validation"`
}

type Request struct {
	Body            any              `json:"body"`
	QueryParameters []QueryParameter `json:"query_parameters"`
	Validations     []BodyVariable   `json:"body_validations"`
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
