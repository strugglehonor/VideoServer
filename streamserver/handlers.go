package main

import (
	"fmt"
	"os"
	"time"
	"net/http"
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