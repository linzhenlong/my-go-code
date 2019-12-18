package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("http get error=",err)
		return
	}
	fmt.Println(res.Cookies())
	data , err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll error=",err)
		return
	}
	fmt.Println(string(data))
}
