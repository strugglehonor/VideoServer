package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/video_server/api/dbops"
	"github.com/video_server/session"
)

var (
	HeaderFieldSessionID = "X-Session-id"
	HeaderFieldUser = "X-User"
)

// validateUserSession
func ValidateUserSession(r *http.Request) {
	username := r.Header[HeaderFieldUser]
	if len(username) == 0 {
		log.Fatal("HTTP header don't have HeaderFieldUser")
	}

	ok, err := session.IsSessionExpired(username[0])
	if err != nil {
		log.Fatal(err.Error())
	}
	if ok {
		log.Fatal("%s's Session has expired", username[0])
	}
}

// validate session id
func ValidateSessionID(r *http.Request) {
	sid := r.Header[HeaderFieldSessionID]
	if len(sid) == 0 {
		log.Fatal("HTTP header don't have HeaderFieldSessionID")
	}

	_, err := dbops.RetrieveSession(sid[0])
	if err != nil {
		log.Fatal(err.Error())
	}

}