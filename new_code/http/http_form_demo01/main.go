package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const form  = "<form action=\"\" method=\"POST\"><p>First name: <input type=\"text\" name=\"fname\" /></p><p>Last name: <input type=\"text\" name=\"lname\" /></p><input type=\"submit\" value=\"Submit\" /></form>"

func test1(writer http.ResponseWriter, request *http.Request)  {
	panic("测试一下异常")
	writer.Header().Set("content-type", "text/html; Charset=utf-8")
	writer.WriteHeader(200)
	fmt.Fprintln(writer,"test1")

}

func test2(writer http.ResponseWriter, request *http.Request)  {
	switch request.Method {
	case http.MethodGet:
		writer.Header().Set("content-type", "text/html; Charset=utf-8")
		writer.WriteHeader(201)
		fmt.Fprintln(writer,form)
	case http.MethodPost:
		request.ParseForm()
		io.WriteString(writer,request.Form["fname"][0])
		io.WriteString(writer,"\n")
		io.WriteString(writer,request.Form["lname"][0])
		io.WriteString(writer,"\n")
		io.WriteString(writer,request.FormValue("fname"))
	}
}

// 处理异常回调
func logPanic(handler http.HandlerFunc)http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if x:=recover();x!=nil{
				log.Printf("[%v] caught panic:%v",request.RemoteAddr, x)
			}
		}()
		handler(writer, request)
	}
	
}

func main() {

	http.HandleFunc("/test1",logPanic(test1))
	http.HandleFunc("/test3",test1)

	http.HandleFunc("/test2", logPanic(test2))
	
	err := http.ListenAndServe("0.0.0.0:8889",nil)
	if err != nil {
		fmt.Println(" http.ListenAndServe(\"0.0.0.0:8889\",nil) error=",err)
	}
}