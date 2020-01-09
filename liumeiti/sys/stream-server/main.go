package main

import (
	"log"
	"github.com/julienschmidt/httprouter"
	"net/http"
)


type middleWareHandler struct {
	r *httprouter.Router
	l *ConnLimiter
}
func NewMiddleWareHandler(r *httprouter.Router, cc int) http.Handler {
	m := middleWareHandler{}
	m.r = r
	m.l = NewConnLimiter(cc)
	return m
}

func (m middleWareHandler)ServeHTTP(w http.ResponseWriter,r *http.Request) {
	if !m.l.GetConn() {
		sendErrorResponse(w, http.StatusTooManyRequests,"Too Many Requests....")
		return
	}
	m.r.ServeHTTP(w, r)
	defer m.l.ReleaseConn()
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/videos/:vid-id", streamHandler)
	router.POST("/upload/:vid-id", uploadHandler)
	router.GET("/testpage",testPage)
	return router
} 

func main() {
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r, 2)
	log.Println("监听8888端口....")
	err := http.ListenAndServe(":8888", mh)
	if err != nil {
		log.Fatalf("http.ListenAndServe err=%s",err.Error())
	}
	
}
