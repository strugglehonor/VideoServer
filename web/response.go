package web

import (
	"io"
	"net/http"
)

// send Normal response
func SendNormalResponse(w http.ResponseWriter, resp string, rc int) {
	w.WriteHeader(rc)
	io.WriteString(w, resp)
}

// send Error Resposne
func SendErrorResponse(w http.ResponseWriter, errMsg string, rc int) {
	w.WriteHeader(rc)
	io.WriteString(w, errMsg)
}
