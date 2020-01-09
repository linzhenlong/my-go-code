package main

import "net/http"

import "github.com/julienschmidt/httprouter"

import "os"

import "time"

import "io/ioutil"

import "log"

import "io"

import "html/template"


func testPage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, err := template.ParseFiles("./html/index.html")
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	shijianchu := time.Now().Unix()
	t.Execute(w,shijianchu)
}

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
	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
	if err := r.ParseMultipartForm(maxUploadSize);err !=nil {
		sendErrorResponse(w, http.StatusBadRequest,"too big form data")
		return
	}
	file, _, err := r.FormFile("file")
	if err !=nil {
		sendErrorResponse(w,http.StatusInternalServerError, err.Error())
		return
	}
	data, err :=ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file error :%v",err)
		sendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	fn := p.ByName("vid-id")
	err = ioutil.WriteFile(videoDir+fn, data, 0666)
	if err != nil {
		log.Printf("WriteFile error:=%v\n",err.Error())
		sendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "upload succ")
}


