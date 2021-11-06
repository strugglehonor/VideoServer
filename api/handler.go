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

// registerUser
func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	username, passwd := getUserInfo(w, r)
	if username != "" && passwd != "" {
		SendErrorResponse(w, defs.RequestParamError)
	}

	err := dbops.AddUserCredential(username, passwd)
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
	// switch to user page after register successfully
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

func LoginHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	username, passwd := getUserInfo(w, r)
	if username != "" && passwd != "" {
		SendErrorResponse(w, defs.RequestParamError)
	}
 
	// true passwd
	tPasswd, err := dbops.GetUserCredential(username)
	if err != nil {
		SendErrorResponse(w, defs.DBSelectError)
	}
	if tPasswd != passwd {
		SendErrorResponse(w, defs.AuthenticateError)
	}

	http.Redirect(w, r, ":8080/UserHome", http.StatusFound)
}

func ListCommentHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid")
	comments, err := dbops.ListCommentPaginate(vid, "", 100, 50)
	if err != nil {
		SendErrorResponse(w, defs.DBSelectError)
	}

	res, _ := json.Marshal(comments)
	SendNormalResponse(w, string(res), http.StatusOK)
}

// postComment
func CreateCommentHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid")

	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		SendErrorResponse(w, defs.RequestParamError)
	}
	
	ubody := &defs.CommentReqs{}
	if err = json.Unmarshal(res, ubody); err != nil {
		SendErrorResponse(w, defs.MarshalError)
	}

	err = dbops.NewComment(vid, ubody.AuthorID, ubody.Content)
	if err != nil {
		SendErrorResponse(w, defs.DBInsertError)
	}

	SendNormalResponse(w, ubody.Content, http.StatusCreated)
}

// listAllVideos
func ListVideoHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("uname")
	videos, err := dbops.GetVideoPaginate(uname, 50, 1000)
	if err != nil {
		SendErrorResponse(w, defs.DBSelectError)
	}

	videosJson, err := json.Marshal(videos)
	if err != nil {
		SendErrorResponse(w, defs.MarshalError)
	}
	SendNormalResponse(w, string(videosJson), http.StatusOK)
}

// createVideo
func CreateVideoHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("uname")
	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		SendErrorResponse(w, defs.RequestParamError)
	}

	videoreq := &defs.VideoReqs{}
	if err = json.Unmarshal(res, videoreq); err != nil {
		SendErrorResponse(w, defs.UnmarshalError)
	}


	video, err := dbops.AddNewVideo(uname, videoreq.VideoName)
	if err != nil {
		SendErrorResponse(w, defs.DBInsertError)
	}
	
	videoJson, err := json.Marshal(video)
	SendNormalResponse(w, string(videoJson), http.StatusCreated)
}

// selectVideo
func GetVideoInfoHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

// get user password and username
func getUserInfo(w http.ResponseWriter, r *http.Request) (string, string) {
	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		SendErrorResponse(w, defs.RequestParamError)
		return "", ""
	}

	ubody := &defs.UserCredential{}
	if err = json.Unmarshal(res, ubody); err != nil {
		SendErrorResponse(w, defs.UnmarshalError)
		return "", ""
	}

	return ubody.Username, ubody.Password
}