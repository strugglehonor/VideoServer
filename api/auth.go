package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/video_server/api/dbops"
	"github.com/video_server/api/session"
)

var (
	HeaderFieldSessionID = "X-Session-id"
	HeaderFieldUser = "X-User"
)

// validateUserSession
func ValidateUserSession(r *http.Request) {
	username := r.Header.Get(HeaderFieldUser)
	if len(username) == 0 {
		log.Fatal("HTTP header don't have HeaderFieldUser")
	}

	ok, err := session.IsSessionExpired(username)
	if err != nil {
		log.Fatal(err.Error())
	}
	if ok {
		log.Fatal("%s's Session has expired", username)
	}
}

// validate session id
func ValidateSessionID(r *http.Request) {
	sid := r.Header.Get(HeaderFieldSessionID)
	if len(sid) == 0 {
		log.Fatal("HTTP header don't have HeaderFieldSessionID")
	}

	_, err := dbops.RetrieveSession(sid)
	if err != nil {
		log.Fatal(err.Error())
	}

}