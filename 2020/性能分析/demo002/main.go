package main


//https://blog.wolfogre.com/posts/go-ppof-practice/
// https://blog.csdn.net/sunxianghuang/article/details/93869683

import (
	"log"

"runtime"

"net/http"

_"net/http/pprof"
)

var datas []string

// Add ...
func Add(str string)string {
	data := []byte(str)
	sData := string(data)
	datas = append(datas,sData)
	return sData
}

func main() {
	if true {
		go func (){
			for {
				log.Println(Add("http://www.baidu.com/"))
			}
		}()
		runtime.GOMAXPROCS(1) // 限制cpu使用数，避免过载
		runtime.SetMutexProfileFraction(1) // 开启对锁调用的跟踪
		runtime.SetBlockProfileRate(1) // 开启对阻塞操作的跟踪
		http.ListenAndServe(":6060",nil)
	}

}