package web

type ApiDef struct {
	Method string  `json:"method"`
	URL    string  `json:"url"`
}

type UserInfo struct {
	Username string `json:"username"`
}