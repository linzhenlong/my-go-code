package main

import (
	"github.com/linzhenlong/my-go-code/2020/study/retriever/real"
	"time"
	"fmt"

	"github.com/linzhenlong/my-go-code/2020/study/retriever/mock"
)

type Retriever interface {
	Get(url string) string
}

func dowload(r Retriever) string  {
	return r.Get("http://www.baidu.com")
}
type Poster interface {
	Post(url string, form map[string]string) string
}
func post(poster Poster) {
	poster.Post("http://www.baidu.com",map[string]string{
		"keyword":"手机",
	})
}

type RetrieverPoster interface {
	Retriever
	Poster
	Connect(host string) 
}
const url = "http://www.baidu.com"
func session(rp RetrieverPoster) string {
	rp.Post(url,map[string]string{
		"Contents" : "test session",
	})
	return rp.Get(url)
}

func inspect(r Retriever) {
	fmt.Println("inspecting", r)
	fmt.Printf("> %T,%v\n", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		 fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:",v.UserAgent)
	}
	fmt.Println()
}
func main() {
	var r Retriever
	r = &mock.Retriever{
		Contents: "this is implement interface",
	}
	inspect(r)
	r = &real.Retriever{
		UserAgent: "php",
		TimeOut: time.Minute,
	}
	inspect(r)
	// 类型断言.
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut)
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}
	//fmt.Println(dowload(r))
	fmt.Println("try a session~~~")
	retriever := &mock.Retriever{
		Contents: "this is a test~",
	}
	fmt.Println(session(retriever))
}