package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/video_server/scheduler/dbops"
)

// del video
func DeleteVideoHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	videoName := p.ByName("video_name")
	err := dbops.AddVideoDeleteRecord(videoName)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, err)
	}
	SendNormalResponse(w)
}
