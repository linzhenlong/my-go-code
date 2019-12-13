package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type User struct {
	Uid string `json:"uid"`
	Name string `json:"name"`
} 

func main()  {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("hello world")
		fmt.Fprintln(writer, "hello1111")
	})

	http.HandleFunc("/user/login", func(writer http.ResponseWriter, request *http.Request) {
		msg := make(map[string]interface{})
		msg["error_code"] = 0
		msg["error_msg"] = "success";
		writer.Header().Set("Content-Type", "application/json;charset=utf-8")
		s, _ := json.Marshal(msg)
		fmt.Println(request.RemoteAddr)
		fmt.Println(request.Host)
		fmt.Println(request.RequestURI)
		fmt.Println(request.UserAgent())
		fmt.Println(request.URL.Query())
		query, _ := url.ParseQuery(request.URL.RawQuery)
		if query == nil {
			msg["error_code"] = 1
			msg["error_msg"] = "userid 不能为空"
		}
		userid := query["user_id"][0]
		username := query["user_name"][0]
		user := User{
			Uid:userid,
			Name:username,
		}
		msg["data"] = user

		writer.Write(s)

	})

	err := http.ListenAndServe("0.0.0.0:8889",nil)
	if err != nil {
		fmt.Println("http.ListenAndServe", err)
	}
}


