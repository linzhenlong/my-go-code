package main

import "net/http"

import "fmt"

// handler函数
func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr,"链接成功")
	fmt.Println(r.Method)
	fmt.Println(r.URL)
	fmt.Println("UA:",r.UserAgent())
	fmt.Println("header:",r.Header)
	fmt.Println("body",r.Body)
	w.Header().Set("Content-Type","application/json")
	w.Write([]byte("hello"))
}

func main() {
	// 回调函数
	http.HandleFunc("/go",myHandler)


	err := http.ListenAndServe(":6060",nil)
	if err !=nil {
		panic(err)
	}
}