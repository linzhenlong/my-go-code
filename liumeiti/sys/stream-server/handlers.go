package main

import "net/http"

import "github.com/julienschmidt/httprouter"

import "os"

import "time"

func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	videoLink := videoDir + vid	
	video, err := os.Open(videoLink)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Status Internal Server Error")
		return
	}
	w.Header().Set("Content-Type", "video/mp4", )
	http.ServeContent(w, r, "", time.Now(), video)
	defer video.Close()
}

func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}


