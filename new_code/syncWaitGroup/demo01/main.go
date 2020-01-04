package main

import "sync"
import "net/http"
import "fmt"

//https://go-zh.org/pkg/sync/#example_WaitGroup

var wg sync.WaitGroup

var urls = []string{
	"http://www.baidu.com",
	"http://www.qq.com",
	"http://www.abc.com",
	"http://www.tmall.com",
}

func main() {

	for _, url := range urls {
		// 递增WaitGroup 计数器
		wg.Add(1)
		// 启动一个goroutine来取回URL
		go func(url string) {
			// 当goroutine 执行完成，减小计数器
			defer wg.Done()

			response,err := http.Get(url)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(url,response.StatusCode)
		}(url)
	}
	// 等待所有的HTTP取回操作.
	wg.Wait()
}