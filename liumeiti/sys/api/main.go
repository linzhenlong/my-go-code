package main

import (
	_ "flag"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

type middleWareHandler struct {
	r  *httprouter.Router
}

// NewMiddleWareHandler 实例化一个中间件handler.
func NewMiddleWareHandler(r *httprouter.Router) http.Handler{
	m := middleWareHandler{}
	m.r = r
	return m
}
// 中间件注入到ServeHTTP 方法中.
func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 检查session.
	validateUserSession(r)
	m.r.ServeHTTP(w, r)
}

var MyLog = &logrus.Logger{}

func init()  {
	MyLog.Out = os.Stdout
	MyLog.Level = logrus.DebugLevel
	MyLog.Formatter = &logrus.TextFormatter{
		ForceColors:true,
		FullTimestamp:true,
		TimestampFormat:"2006-01-02 15:04:05.000",
	}
}



func UserHandlers() *httprouter.Router {
	router := httprouter.New()
	
	router.POST("/user", CreateUser)

	router.POST("/user/:user_name", Login)
	return router
}

func main() {

	/*logPath := flag.String("logPath", "/tmp/limeiti.log", "日志路径")

	flag.Parse()

	fd , err := os.OpenFile(*logPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0655)

	if err == nil {
		log.Out = fd
		defer fd.Close()
	}*/

	r := UserHandlers()
	middleWareHandler := NewMiddleWareHandler(r)

	MyLog.Infof("监听8889端口")
	err := http.ListenAndServe("0.0.0.0:8889",middleWareHandler)
	if err !=nil {
		MyLog.Errorf("http.ListenAndServe 8889 端口 error= ", err)
	}
}
