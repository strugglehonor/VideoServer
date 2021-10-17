package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// register
func RegisterHandler() *httprouter.Router {
	router := httprouter.New()

	router.POST("/delVideo/:video_name", DeleteVideoHandler)

	return router
}

func main() {
	r := RegisterHandler()
	http.ListenAndServe(":9001", r)
}
