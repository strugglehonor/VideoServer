package main

import (
	"io"
	"net/http"
)

// send response
func SendNormalResponse(w http.ResponseWriter) {
	io.WriteString(w, "ok")
}

// send error response
func SendErrorResponse(w http.ResponseWriter, sc int, err error) {
	w.WriteHeader(sc)
	io.WriteString(w, err.Error())
}
