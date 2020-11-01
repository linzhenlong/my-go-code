package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/BurntSushi/toml"
	"github.com/linzhenlong/my-go-code/2020/网络/proxy/config"
	"github.com/linzhenlong/my-go-code/2020/网络/proxy/util"
)

// MyProxyHandler ...
type MyProxyHandler struct {
}

func (m *MyProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("r.RequestURI=>", r.RequestURI)
	//fmt.Println("r.URL.Path=>", r.URL.Path)
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(500)
			log.Println(err)
		}
	}()

	// 代理规则
	// 排除/favicon.ico 的请求
	if r.URL.Path != "/favicon.ico" {
		//url, _ := url.Parse(lb.SelectByRand().Host)
		//url, _ := url.Parse(util.LB.SelectByIPHash(r.RemoteAddr).Host)
		url, _ := url.Parse(util.LB.RoundRobin().Host)
		reverseProxy := httputil.NewSingleHostReverseProxy(url)
		reverseProxy.ServeHTTP(w, r)
	}

	/**
	if len(cfg.WebList.Web) > 0 {
		for _, val := range cfg.WebList.Web {
			if ok, _ := regexp.MatchString(val.Path, r.URL.Path); ok {
				log.Println(r.RemoteAddr)
				// 用户访问http://localhost:8080/a 时代理到9091的web上去

				// 通过httpclient 实现反向代理
				//util.RequestURL(w, r, val.Pass)

				// 使用go内置的反向代理
				target, _ := url.Parse(val.Pass + val.Path)
				reverseProxy := httputil.NewSingleHostReverseProxy(target)
				// 自定义一下超时时间，
				myTransport := &http.Transport{
					DialContext: (&net.Dialer{
						Timeout:   30 * time.Second, // tcp链接的超时时间.
						KeepAlive: 30 * time.Second,
					}).DialContext,
					ResponseHeaderTimeout: 10 * time.Second, // 读取header的超时时间.
				}
				reverseProxy.Transport = myTransport
				reverseProxy.ServeHTTP(w, r)
				return
			}
		}
	}
	**/
	//w.Write([]byte("proxy index"))
}

var (
	// Cfg ...
	Cfg config.TomlConfig
)

func main() {

	// 加载配置文件
	_, err := toml.DecodeFile("../env.toml", &Cfg)
	if err != nil {
		log.Println("加载配置文件出错:", err)
		return
	}
	http.ListenAndServe(":8081", &MyProxyHandler{})

}
