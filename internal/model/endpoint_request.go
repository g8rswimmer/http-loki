package model

type Request struct {
	Body any `json:"body"`
}

type Response struct {
	StatusCode int `json:"status_code"`
	Body       any `json:"body"`
}

type Mock struct {
	Method   string   `json:"method"`
	Endpoint string   `json:"endpoint"`
	Request  Request  `json:"request"`
	Response Response `json:"response"`
}
