package main

import "context"

import "time"

import "net/http"

import "fmt"

import "sync"

import "io/ioutil"


type respData struct {
	resp *http.Response
	err error
}

func doCall(ctx context.Context) {
	transport := http.Transport{
		// 请求频繁可定义全局的Client对象并开启长连接
		// 请求不频繁使用短连接
		DisableKeepAlives: true,
	}
	client := http.Client{
		Transport: &transport,
	}
	respChan := make(chan *respData,1)
	req,err := http.NewRequest("GET","http://127.0.0.1:8889", nil)
	if err != nil {
		fmt.Printf("new requestg failed, err:%v\n", err)
		return
	}
	req = req.WithContext(ctx) // 使用带超时的ctx创建一个新的client request

	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Wait()
	go func(){
		resp , err := client.Do(req)
		fmt.Printf("client DO resp:%v,err:%v\n",resp, err)
		rd := &respData{
			resp: resp,
			err: err,
		}
		respChan<- rd
		wg.Done()
	}()
	select{
	case <-ctx.Done():
		fmt.Println("call api timeout")
	case result := <- respChan:
		fmt.Println("call api succ")
		if result.err != nil {
			fmt.Printf("call server api failed, err:%v\n", result.err)
			return
		}
		defer result.resp.Body.Close()
		data,_ := ioutil.ReadAll(result.resp.Body)
		fmt.Printf("resp:%v\n",string(data))
	}
}

func main() {
	// 定义一个100ms超时的上下文
	ctx,cancel := context.WithTimeout(context.Background(),time.Millisecond * 100)

	defer cancel()

	// 请求服务端.
	doCall(ctx)
}