package web

type ApiDef struct {
	Method  string  `json:"method"`
	URL     string  `json:"url"`
	ReqBody string	`json:"req_body"`
}