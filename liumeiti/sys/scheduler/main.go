package main
import (
	"github.com/linzhenlong/my-go-code/liumeiti/sys/scheduler/taskrunner"
	"github.com/julienschmidt/httprouter"
	"net/http"
	
)
func RegisterHandles() *httprouter.Router {
	router := httprouter.New()
	router.GET("/video-delete-record/:vid-id", videoDelRecHandler)
	return router
}


func main() {
	go taskrunner.Start()
	r := RegisterHandles()
	http.ListenAndServe(":8887",r)
}