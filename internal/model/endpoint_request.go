package model

type VariableParams struct {
	Args   []string `json:"args"`
	Prefix string   `json:"prefix"`
	Suffix string   `json:"suffix"`
}

type BodyVariable struct {
	VariableParams
	Path string `json:"json_path"`
	Func string `json:"func"`
}

type QueryVariable struct {
	VariableParams
	Func string `json:"func"`
}

type QueryParameter struct {
	Key        string         `json:"key"`
	Value      string         `json:"value"`
	Validation *QueryVariable `json:"validation"`
}

type Request struct {
	Body            any              `json:"body"`
	QueryParameters []QueryParameter `json:"query_parameters"`
	Validations     []BodyVariable   `json:"body_validations"`
}

func (r Request) QueryVariables() []QueryVariable {
	vars := make([]QueryVariable, 0, len(r.QueryParameters))
	for _, q := range r.QueryParameters {
		if q.Validation == nil {
			continue
		}
		vars = append(vars, *q.Validation)
	}
	return vars
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
