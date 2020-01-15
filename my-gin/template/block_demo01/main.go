package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 定义模板
	// 解析模板
	t, err := template.ParseFiles("./index.tmpl")
	if err != nil {
		log.Printf("解析模板错误:%s", err.Error())
	}
	// 渲染模板
	msg := "丰台小王子"
	t.Execute(w,msg)
}
func index2(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 定义模板 模板继承的方式
	// 解析模板, 需要根模板写在前面，也就是先加载base模板，index是继承base 所以要放在后面
	t, err := template.ParseFiles("./templates/base.tmpl","./templates/index.tmpl")
	if err != nil {
		log.Printf("解析模板错误:%s", err.Error())
	}
	// 渲染模板
	msg := "我是继承base模板的丰台小王子"
	// ExecuteTemplate 需要指定渲染那个模板
	t.ExecuteTemplate(w,"index.tmpl",msg)
}
func home(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 定义模板 
	// 解析模板
	t, err := template.ParseFiles("./home.tmpl")
	if err != nil {
		log.Printf("解析模板错误:%s", err.Error())
	}
	// 渲染模板
	msg := "丰台小王子"
	t.Execute(w,msg)
}
func home2(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 定义模板 模板继承的方式
	// 解析模板, 需要根模板写在前面，也就是先加载base模板，index是继承base 所以要放在后面
	t, err := template.ParseFiles("./templates/base.tmpl","./templates/home.tmpl")
	if err != nil {
		log.Printf("解析模板错误:%s", err.Error())
	}
	// 渲染模板
	msg := "我是继承base模板的丰台小王子home2"
	// ExecuteTemplate 需要指定渲染那个模板
	t.ExecuteTemplate(w,"home.tmpl",msg)
}

func logPanic(handle httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer func() {
			if x := recover();x != nil {
				log.Printf("catch panic :%v",x)
			}
		}()
		handle(w, r, p)
	}
}
func registerHTTPHandlers() *httprouter.Router{
	r := httprouter.New()
	r.GET("/index",logPanic(index))
	r.GET("/index2",logPanic(index2))
	r.GET("/home",logPanic(home))
	r.GET("/home2",logPanic(home2))
	return r
}

func main() {
	router := registerHTTPHandlers()
	http.ListenAndServe(":9000", router)
}
//https://www.bilibili.com/video/av79417575/