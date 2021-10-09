package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func RegisterHander() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:username", Login)
	// httprouter支持对特殊情况下的定制
	// router.NotFound = http.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Not Found!"))
	// })
	return router
}

func main() {
	r := RegisterHander()
	http.ListenAndServe(":8000", r)
}