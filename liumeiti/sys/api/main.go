package main

import (
	_ "flag"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)



var log = &logrus.Logger{}

func init()  {
	log.Out = os.Stdout
	log.Level = logrus.DebugLevel
	log.Formatter = &logrus.TextFormatter{
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


	log.Infof("监听8889端口")
	err := http.ListenAndServe("0.0.0.0:8889",r)
	if err !=nil {
		log.Errorf("http.ListenAndServe 8889 端口 error= ", err)
	}
}
