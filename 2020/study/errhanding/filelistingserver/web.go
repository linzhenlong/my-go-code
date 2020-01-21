package main

import (
	"github.com/linzhenlong/my-go-code/2020/study/errhanding/filelistingserver/filelisting"
	"log"
	_"github.com/gpmgo/gopm/log"
	"net/http"
	"os"
)

type appHandler func(w http.ResponseWriter, r *http.Request) error

func errWrapper(handler appHandler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		code := http.StatusOK
		if err != nil {
			//log.("Eroor handing request ")
			//log.Warn("Eroor handing request:%s",err.Error())
			log.Printf("Eroor handing request:%s", err.Error())
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(w, http.StatusText(code), code)
		}
	}
}

func main() {
	http.HandleFunc("/list/", errWrapper(filelisting.HandleFileList))
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		panic(err)
	}
}

// 7-4