package web

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()

	router.GET("/UserHome", UserHomeHandler)

	return router
}

func main() {
	r := RegisterHandler()
	http.ListenAndServe(":8080", r)
}

