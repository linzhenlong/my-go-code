package main

import (
	"github.com/julienschmidt/httprouter"
	"time"
	"math/rand"
	"strconv"
	"html/template"
	"log"
	"io/ioutil"
	"net/http"
)
// User 结构体.
type User struct {
	ID int
	Name string
	Age int
	Gender string
}

func getAge(max ,min int) int {
	if min > max {
		return min
	}
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(max)
	if randNum > min {
		return randNum
	}
	if randNum < min {
		return min
	}
	return randNum
}
// FuncHandler handler.
func FuncHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	htmlByte, err := ioutil.ReadFile("./hello.tmpl")
	if err != nil {
		log.Fatalf("ioutil.ReadFile error:%s", err.Error())
	}
	// 自定义一个kua 方法
	kua := func(arg string)(string, error) {
		return arg+"真棒啊~~~",nil
	}
	// 自定义偶数函数
	oushu := func(num int)(int) {
		return num%2
	}
	tmpl, err := template.New("hello").Funcs(template.FuncMap{"kua":kua,"oushu":oushu}).Parse(string(htmlByte))
	if err != nil {
		log.Printf("template 解析失败:%s", err.Error())
	}
	var res []User
	for i:=1;i<=15;i++ {
		user := User{
			ID: i,
			Name: "学生"+strconv.Itoa(i),
			Age: getAge(30,18),
			Gender: "男",
		}
		res = append(res, user)
	}
	tmpl.Execute(w, res)
	
}
// 处理panic 防止由于panic 导致http服务崩溃. 
func logPanic(handle httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request ,p httprouter.Params) {
		defer func() {
			if x:=recover();x !=nil {
				log.Printf("catch panic %v",x)
			}
		}()
		handle(w, r , p)
	}
}
// RegisterHTTPHandlers 注册路由.
func RegisterHTTPHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/func",logPanic(FuncHandler))
	//router.GET("/func",FuncHandler)
	return router
}

func main() {
	r := RegisterHTTPHandlers()
	http.ListenAndServe(":8889", r)
}
//https://www.bilibili.com/video/av78808893/