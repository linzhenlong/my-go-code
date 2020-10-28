package main

import (
	"net/http"
	"fmt"
	"log"
	"io/ioutil"
	"github.com/linzhenlong/my-go-code/2020/网络/proxy/util"
)

type MyProxyHandler struct {

}

func (m *MyProxyHandler)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RequestURI)
	fmt.Println(r.URL.Path)
	defer func(){
		if err := recover();err!=nil {
			w.WriteHeader(500)
			log.Println(err)
		}
	}()
	if r.URL.Path == "/a" {
		// 用户访问http://localhost:8080/a 时代理到9091的web上去
		newRequest, _ := http.NewRequest(r.Method,"http://localhost:9091",r.Body)

		// 将客户端的请求头，拷贝给目标网站的请求头
		util.CloneHeader(r.Header,&newRequest.Header)
		// 发送请求
		newResp, _:= http.DefaultClient.Do(newRequest)
		
		getHeader :=w.Header()
		// 将目标网站的响应头拷贝给 客户端
		util.CloneHeader(newResp.Header,&getHeader)
		w.WriteHeader(newResp.StatusCode)
		defer newResp.Body.Close()
		resContentBytes, _ := ioutil.ReadAll(newResp.Body)
		w.Write(resContentBytes)
		return
	}
	w.Write([]byte("proxy index"))
}

func main() {
	http.ListenAndServe(":8081",&MyProxyHandler{})

}

6让我们的 反向代理 支持Basic Auth验证、获取真实IP【瑞客论坛 www.ruike1.com 的   ======》第2分钟