package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
)

type web1Handler struct {
}
type web2Handler struct {
}
type web3Handler struct {
}

// GetIP ...
func (w *web1Handler) GetIP(r *http.Request) string {
	ips := r.Header.Get("x-forwarded-for")
	if ips != "" {
		ipsList := strings.Split(ips, ",")
		if len(ipsList) > 0 && ipsList[0] != "" {
			return ipsList[0]
		}
	}
	return r.RemoteAddr
}

// ServeHTTP ...
func (w *web1Handler) ServeHTTP(writer http.ResponseWriter, r *http.Request) {
	log.Println("this is web1")
	/**
	// basic auth 认证.
	auth := r.Header.Get("Authorization")
	if auth == "" {
		writer.Header().Set("WWW-Authenticate",`Basic realm="请输入用户及密码"`)
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	authList := strings.Split(auth," ")

	if len(authList) == 2 && authList[0] == "Basic"{
		res, err := base64.StdEncoding.DecodeString(authList[1])
		if err == nil && string(res) == "linzl:123" {
			out := fmt.Sprintf("<h1>web1,来自于:%s</h1>",w.GetIP(r))
			writer.Write([]byte(out))
			return
		}
	}
	**/
	writer.Write([]byte("web1 用户名密码错误"))
}
func (*web2Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//time.Sleep(time.Second * 3)
	log.Println("this is web2")
	w.Write([]byte("web2"))

}
func (*web3Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//time.Sleep(time.Second * 3)
	log.Println("this is web3")
	w.Write([]byte("web3"))

}

// 利用两个协程启动两个web
func main() {
	c := make(chan os.Signal)
	// web1 9091
	go func() {
		http.ListenAndServe(":9091", &web1Handler{})

	}()
	// web2 9092
	go func() {
		http.ListenAndServe(":9092", &web2Handler{})
	}()
	// web3 9093
	go func() {
		http.ListenAndServe(":9093", &web3Handler{})
	}()
	signal.Notify(c, os.Interrupt)
	s := <-c
	log.Println(s)
}
