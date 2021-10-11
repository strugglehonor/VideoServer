package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//middleware check session
	ValidateUserSession(r)
	ValidateSessionID(r)
	m.r.ServeHTTP(w, r)
}

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
	m := NewMiddleWareHandler(r)
	http.ListenAndServe(":8000", m)
}