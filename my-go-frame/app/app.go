package app

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/linzhenlong/my-go-code/my-go-frame/conf"
)

// FrameAPP 框架实例...
type FrameAPP struct {
	Gin *gin.Engine
}

// New 初始化..
func New() (app *FrameAPP, err error) {
	// 加载配置文件
	err = conf.InitConf()
	if err != nil {
		return
	}
	fmt.Printf("xxxx%#v", conf.FrameConf)
	// 创建app
	app = &FrameAPP{}

	// 初始化gin框架
	app.initGin()

	// todo 创建客户端

	// todo 注册中间件
	return
}

// 初始化gin...
func (f *FrameAPP) initGin() {
	if conf.FrameConf.Debug != 0 {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	f.Gin = gin.New()
}

// Run 启动..
func (f *FrameAPP) Run() (err error) {
	srv := &http.Server{
		Handler:      f.Gin,
		ReadTimeout:  time.Duration(conf.FrameConf.AppConf.ReadTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(conf.FrameConf.AppConf.WriteTimeout) * time.Millisecond,
	}
	go func() {
		listener, err := net.Listen("tcp", conf.FrameConf.AppConf.Addr)
		if err != nil {
			errMsg := fmt.Sprintf("启动server失败:%+v,%+v", srv, err)
			panic(errMsg)
		}
		srv.Serve(listener)
	}()
	return
}

func (f *FrameAPP) waitGraceExit(server *http.Server) {

}
