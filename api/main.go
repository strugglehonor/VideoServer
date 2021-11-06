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

	router.GET("/videos/:vid/comments", ListCommentHandler)

	router.GET("/listVideo", ListVideoHandler)

	router.POST("/createComment", CreateCommentHandler)

	router.POST("/createVideo", CreateVideoHandler)

	router.GET("/video/:vid", GetVideoInfoHandler)
	
	return router
}

func main() {
	r := RegisterHandler()
	m := NewMiddleWareHandler(r)
	http.ListenAndServe(":8000", m)
}
