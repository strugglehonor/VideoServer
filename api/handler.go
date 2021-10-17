package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/video_server/api/dbops"
	"github.com/video_server/api/defs"
	"github.com/video_server/api/session"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		SendErrorResponse(w, defs.RequestParamError)
		return
	}

	ubody := &defs.UserCredential{}
	if err = json.Unmarshal(res, ubody); err != nil {
		SendErrorResponse(w, defs.UnmarshalError)
		return
	}

	username, passwd := ubody.Username, ubody.Password
	err = dbops.AddUserCredential(username, passwd)
	if err != nil {
		SendErrorResponse(w, defs.DBInsertError)
		return
	}

	id, err := session.NewSessionID(username)
	if err != nil {
		SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}

	su := &defs.SignedUp{Success: true, SessionID: id} // response
	resp, err := json.Marshal(su)
	if err != nil {
		SendErrorResponse(w, defs.MarshalError)
		return
	}

	SendNormalResponse(w, string(resp), http.StatusCreated)

}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	username := p.ByName("username")
	isExpired, err := session.IsSessionExpired(username)
	if err != nil {
		SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}

	if isExpired {
		SendErrorResponse(w, defs.SessionExpiredError)
		return
	}

	msg := fmt.Sprintf("username:%s login success", username)
	SendNormalResponse(w, msg, http.StatusAccepted)
}
