package main

import (
	"fmt"
	"io"
	"net/http"
)

const form  = "<form action=\"\" method=\"POST\"><p>First name: <input type=\"text\" name=\"fname\" /></p><p>Last name: <input type=\"text\" name=\"lname\" /></p><input type=\"submit\" value=\"Submit\" /></form>"

func main() {


	http.HandleFunc("/test1", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("content-type", "text/html; Charset=utf-8")
		writer.WriteHeader(200)
		fmt.Fprintln(writer,"test1")
	})

	http.HandleFunc("/test2", func(writer http.ResponseWriter, request *http.Request) {
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
	})
	
	err := http.ListenAndServe("0.0.0.0:8889",nil)
	if err != nil {
		fmt.Println(" http.ListenAndServe(\"0.0.0.0:8889\",nil) error=",err)
	}


}