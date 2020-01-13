package main

import "net/http"

import "html/template"

import "log"

type Preson struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func Test(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./hello.tmpl")
	person := Preson{
		Name: "沙河 小王子",
		Age: 18,
	}
	myMap  := make(map[string]Preson)
	myMap["person1"] = person
	
	hobbyList := []string{
		"打篮球",
		"打泰拳",
		"踢足球",
		"xxoo",
	}
	t.Execute(w,map[string]interface{}{
		"person":person,
		"myMap":myMap,
		"hobby":hobbyList,
	})
}

func LogPanic(handler http.HandlerFunc)http.HandlerFunc {
	return func(w http.ResponseWriter,r *http.Request){
		defer func() {
			if x:=recover();x!=nil {
				log.Printf("caught panic:%v",x)
			}
		}()
		handler(w,r)
	}
}

func main() {

	http.HandleFunc("/",LogPanic(Test))
	log.Printf(":8889端口启动....")
	err := http.ListenAndServe(":8889",nil)
	if err !=nil {
		panic("ListenAndServe :8889 error"+err.Error())
	}
}
//https://www.bilibili.com/video/av78808893/?spm_id_from=333.788.videocard.0