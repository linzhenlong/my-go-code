package main

import (
	"html/template"
	"github.com/julienschmidt/httprouter"
	"strconv"
	"math"
	"log"
	"net/http"
)

type Person struct {
	Name string
	Age int
	Adress string
	Hobby []string
}

// bsf 自定义标识符.
func bsf(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	/* hobby := []string {
		"打篮球",
		"踢足球",
		"打泰拳",
		"游泳",
	} */
	t := template.New("base.tmpl")
	avg := func(a,b int) float64 {
		tmep := float64(a / b)
		return math.Round(tmep)
	}
	// Delims 方法使用自定义的标识符
	t, err := t.Funcs(template.FuncMap{"avg":avg}).Delims("[[","]]").ParseFiles("./templates/base.tmpl")
	if err != nil {
		log.Printf("模板解析失败")
	}
	var data []Person
	for i:=0;i<10;i++ {
		person := Person{
			Name: "学生1"+strconv.Itoa(i),
			Age: 18+i,
			Adress: "丰田",
			Hobby: []string{
				"打篮球",
			},
		}
		data = append(data, person) 
	}
	t.Execute(w, data)
	//template.ParseFiles("./templates/base.tmpl")
}

func xss(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	
	// 故意定义一个有风险的字符串
	// 如果导入的text/template包，js代码在渲染之后会被执行.
	// 如果导入的是html/template包,js代码不会执行，会被转实体.
	// 但是有些时候像是"<a href='http://baidu.com'>百度</a>" 不需要转实体
	// 因此需要一个自定义的函数处理一下
	t, err := template.New("xss.tmpl").
		Funcs(template.FuncMap{
			"safe":func(str string) template.HTML{
				return template.HTML(str)
			},
		}).
		Delims("[[","]]").
		ParseFiles("./templates/xss.tmpl")
	if err != nil{
		log.Printf("模板解析失败")
		return
	}
	data := make(map[string]string,2)
	xssStr := "<script>alert('xss');</script>"
	data["xss"] = xssStr
	aStr := "<a href='http://baidu.com'>百度</a>"
	data["a"] = aStr
	err = t.Execute(w, data)
	if err != nil {
		log.Printf("execute err:%s",err.Error())
		return
	}

}
// LogPanic 日志输出panic.
func LogPanic(handler httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer func() {
			if x := recover();x!=nil {
				log.Printf("catch panic:%v", x)
			}
		}()
		handler(w, r, p)
	}
}

func RegisterHandlers() *httprouter.Router {
	r := httprouter.New()
	r.GET("/bsf", LogPanic(bsf))
	r.GET("/xss", LogPanic(xss))
	return r
}

func main() {
	router := RegisterHandlers()
	http.ListenAndServe(":9000", router)
}

//https://www.bilibili.com/video/av79671612/