package main

import (
	"fmt"
	"regexp"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/linzhenlong/my-go-code/2020/网络/proxy/config"
	"github.com/linzhenlong/my-go-code/2020/网络/proxy/util"
)

// MyProxyHandler ...
type MyProxyHandler struct {
}

func (m *MyProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RequestURI)
	fmt.Println(r.URL.Path)
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(500)
			log.Println(err)
		}
	}()
	
	// 代理规则
	if len(cfg.WebList.Web) > 0 {
		for _, val := range cfg.WebList.Web {
			if ok,_:=regexp.MatchString(val.Path,r.URL.Path);ok {
				log.Println(r.RemoteAddr)
				// 用户访问http://localhost:8080/a 时代理到9091的web上去
				util.RequestURL(w, r, val.Pass)
			return
			}
		}
	}
	w.Write([]byte("proxy index"))
}

var (
	cfg config.TomlConfig
)

func main() {

	// 加载配置文件
	_, err := toml.DecodeFile("../env.toml", &cfg)
	if err != nil {
		log.Println("加载配置文件出错:", err)
		return
	}
	
	http.ListenAndServe(":8081", &MyProxyHandler{})

}
