package main

import (
	"fmt"
	"net/http"
	"time"
)

func main()  {

	var url = []string {
		"http://www.baidu.com",
		"http://g.cn",
		"http://taobao.com",
		"http://smzdm.com",
		"http://xxx.lo",
	}

	for _, v := range url {

		client := &http.Client{
			Timeout:time.Second,
		}
		resp,err := client.Head(v)
		if err != nil {
			fmt.Println(v,"http.Head(v) error=",err)
			continue
		}
		fmt.Println(v,"的状态码是:",resp.StatusCode)
	}
}
