package web

import "github.com/julienschmidt/httprouter"

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()

	router.GET("/UserHome", UserHomeHandler)

	return router
}

func main() {

}

