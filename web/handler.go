package web

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/video_server/api/dbops"
	s "github.com/video_server/api/session"
)

// user home render
func UserHomeHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	session, err1 := r.Cookie("session")
	username, err2 := r.Cookie("username")

	if err1 != nil || err2 != nil {
		errMsg := fmt.Sprintf("Error occur when read cookie:%s;\n%s", err1.Error(), err2.Error())
		SendErrorResponse(w, errMsg, http.StatusInternalServerError)
		return
	}

	if len(session.Value) != 0 || len(username.Value) != 0 {
		ok, err1 := s.IsSessionExpired(username.Value)
		ss, err2 := dbops.RetrieveSession(session.Value)

		if err1 != nil || err2 != nil {
			errMsg := fmt.Sprintf("Error occur when check session:%s;\n%s", err1.Error(), err2.Error())
			SendErrorResponse(w, errMsg, http.StatusUnauthorized)

			return
		}

		if ok {
			SendErrorResponse(w, "Login has expired.", http.StatusUnauthorized)
		}

		if ss.UserName != username.Value {
			SendErrorResponse(w, "username conflict", http.StatusInternalServerError)
		}

		u := UserInfo{Username: username.Value}
		err := renderPage(w, "../template/userhome.html", u)
		if err != nil {
			SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		}
		
	}
}

// render home
func renderPage(w http.ResponseWriter, filepath string, data  interface{}) error {
	t, err := template.ParseFiles(filepath)	
	if err != nil {
		return err
	}

	err = t.Execute(w, data)
	if err != nil {
		return err
	}

	return nil
}

// render home page
func HomeHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := renderPage(w, "../template/home.html", "ZigmundSu")
	if err != nil {
		SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		
	}
}