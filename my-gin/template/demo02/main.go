package main

import "net/http"

import "html/template"

import "log"

type Preson struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func Test(w http.ResponseWriter, r *http.Request) {

	/* // 自定义函数 	
	kua := func(arg string)(string, error) {
		return arg+"真漂亮",nil
	} */
	//t, err := template.New("hello").Funcs(template.FuncMap{"kua":kua}).ParseFiles("./hello.tmpl")
	t, err := template.ParseFiles("./hello.tmpl")
	if err !=nil {
		log.Printf("template error%s",err.Error())
	}
	// = template.ParseFiles("./hello.tmpl")
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
