package main

import (
	"io"
	"net/http"
)

func SendErrorResponse(w http.ResponseWriter, errCode int, errMsg string) {
	w.WriteHeader(errCode) // statuscode
	// w.Write([]byte(errMsg)) may ok,too
	io.WriteString(w, errMsg)
}