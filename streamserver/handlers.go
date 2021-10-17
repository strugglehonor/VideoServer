package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/julienschmidt/httprouter"
)

// play video
func StreamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("video-id")
	filePath := fmt.Sprintf("%s%s", VideoDir, vid)

	video, err := os.Open(filePath)
	if err != nil {
		errMsg := fmt.Sprintf("video:%s%s open failed", VideoDir, video.Name())
		SendErrorResponse(w, 500, errMsg)
	}

	w.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(w, r, "", time.Now(), video)

	defer video.Close()
}

// upload video
func UploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// request's size is limited
	r.Body = http.MaxBytesReader(w, r.Body, MaxUploadSize)
	if err := r.ParseMultipartForm(MaxUploadSize); err != nil {
		SendErrorResponse(w, http.StatusBadRequest, err.Error())
		log.Error(err)
		return
	}

	file, _, err := r.FormFile("file") // In frontend, key must be file.
	if err != nil {
		log.Error("file open failed: %v", err)
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Error("Error occur when get file data: %d", err)
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	fn := p.ByName("video-id")
	if err = ioutil.WriteFile(VideoDir+fn, data, 0666); err != nil {
		log.Error("Write file Error: %v", err)
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "upload successfully")
}

// test page
func TestPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, _ := template.ParseFiles("/video/form.html")
	t.Execute(w, nil)
}
