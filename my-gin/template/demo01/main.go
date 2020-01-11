package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		log.Printf("解析模板错误%s\n", err.Error())
	}
	// 渲染模板
	t.Execute(w, "沙河小王子")
}

func logPanic(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("caught panic:%v", r.RemoteAddr)
			}
		}()
		handler(w, r)
	}
}
func main() {
	http.HandleFunc("/hello", logPanic(sayHello))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Printf(":8888 启动失败%s\n", err.Error())
	}

}

//https://www.bilibili.com/video/av78578060
// 【Go Web开发系列教程】04-Go template模板初识
