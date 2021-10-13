package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type middleWareHandler struct {
	r *httprouter.Router
	c *ConnLimiter
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	m.c.bucket = make(chan int, limitVal)
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.c.GetConnLimiter()
	m.r.ServeHTTP(w, r)
}

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()
	router.GET("/stream/:video-id", StreamHandler)
	return router
}

func main() {
	r := RegisterHandler()
	http.ListenAndServe(":9000", r)
}
