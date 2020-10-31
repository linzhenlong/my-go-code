package util

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

// CloneHeader 拷贝http header
func CloneHeader(src http.Header, dest *http.Header) {
	for k, v := range src {
		dest.Set(k, v[0])
	}
}

// RequestURL ...
func RequestURL(writer http.ResponseWriter, request *http.Request, url string) {
	// 创建一个新的request 请求
	newRequest, _ := http.NewRequest(request.Method, url, request.Body)

	// 拷贝当前请求的header到新的http请求
	CloneHeader(request.Header, &newRequest.Header)

	// 设置 x-forwarded-for
	newRequest.Header.Set("x-forwarded-for", request.RemoteAddr)

	// 自定义一下超时时间，
	myTransport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second, // tcp链接的超时时间.
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ResponseHeaderTimeout: 1 * time.Second, // 读取header的超时时间.
	}

	//执行request请求
	//newResponse, _:=http.DefaultClient.Do(newRequest)
	newResponse, err := myTransport.RoundTrip(newRequest)
	if err != nil {
		log.Println("myTransport.RoundTrip err:",err)
	}
	// 拷贝响应头给客户端
	newRespHeader := writer.Header()
	CloneHeader(newResponse.Header, &newRespHeader)
	writer.WriteHeader(newResponse.StatusCode)

	defer newResponse.Body.Close()
	resp, _ := ioutil.ReadAll(newResponse.Body)
	writer.Write(resp)

}
