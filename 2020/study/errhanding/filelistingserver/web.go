package main

import (
	"github.com/linzhenlong/my-go-code/2020/study/errhanding/filelistingserver/filelisting"
	"log"
	"net/http"
	"os"
	_ "net/http/pprof" //http://127.0.0.1:9000/debug/pprof/ 性能分享.
)
type userError interface{
	error
	Message() string
}

type appHandler func(w http.ResponseWriter, r *http.Request) error

func errWrapper(handler appHandler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic:%v", r)
				http.Error(w, "服务器开小差了...", http.StatusInternalServerError)
			}
		}()
		err := handler(w, r)
		code := http.StatusOK
		if err != nil {
			//log.("Eroor handing request ")
			//log.Warn("Eroor handing request:%s",err.Error())
			log.Printf("Eroor handing request:%s", err.Error())
			if userErr, ok := err.(userError);ok {
				http.Error(w, userErr.Message(), http.StatusBadRequest)
				return
			}
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
	http.HandleFunc("/", errWrapper(filelisting.HandleFileList))
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		panic(err)
	}
}

// 7-4
