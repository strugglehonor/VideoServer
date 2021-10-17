package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
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
	// ValidateUserSession(r)
	// ValidateSessionID(r)
	m.r.ServeHTTP(w, r)
}

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:username", Login)
	return router
}

func main() {
	r := RegisterHandler()
	m := NewMiddleWareHandler(r)
	http.ListenAndServe(":8000", m)
}
