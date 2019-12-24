package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Person struct {
	Name string
	Age int
}

func test(response http.ResponseWriter,request *http.Request)  {
	//
	p := Person{
		Name:"tom",
		Age:18,
	}
	t, err := template.ParseFiles("./index.html")
	if err != nil {
		fmt.Fprintln(os.Stderr, "template.ParseFiles error=", err)
	}
	t.Execute(response,p)
}

func logPanic(handler http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if err:=recover();err !=nil {
				log.Printf("[%v] caught panic:%v",request.RemoteAddr, err)
			}
		}()
		handler(writer, request)
	}
}

func main() {

	http.HandleFunc("/test",logPanic(test))

	err := http.ListenAndServe("0.0.0.0:8889",nil)
	if err != nil {
		fmt.Println("http.ListenAndServe error=", err)
		return
	}
}
