package main

import (
	"encoding/json"
	"github.com/video_server/api/defs"
	"io"
	"net/http"
)

// send Normal response
func SendNormalResponse(w http.ResponseWriter, resp string, rc int) {
	w.WriteHeader(rc)
	io.WriteString(w, resp)
}

// send Error Resposne
func SendErrorResponse(w http.ResponseWriter, err defs.Error) {
	w.WriteHeader(err.HttpCode)
	resStr, _ := json.Marshal(&err.ErrorDetail)
	io.WriteString(w, string(resStr))
}
