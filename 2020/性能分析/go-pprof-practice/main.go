package main

import (
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
	"github.com/linzhenlong/my-go-code/2020/性能分析/go-pprof-practice/animal"
	_ "net/http/pprof"
)

func main() {

	// 日志
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.SetOutput(os.Stdout)

	runtime.GOMAXPROCS(1)              // 限制cpu使用数，避免过载
	runtime.SetMutexProfileFraction(1) // 开启对锁调用的跟踪
	runtime.SetBlockProfileRate(1)     // 开启对阻塞操作的跟踪

	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	for {
		for _,v := range animal.Allanimals {
			v.Live()
		}
		
		time.Sleep(time.Second)
	}

}
