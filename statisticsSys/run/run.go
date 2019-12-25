package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type resource struct {
	url string
	target string
	start int
	end int
}

func ruleResource()[]resource  {
	var res []resource
	r1 := resource{
		url:"http://gxcms.lo/",
		target:"",
		start:0,
		end:0,
	}
	r2 := resource{
		url:"http://gxcms.lo/{$id}/",
		target:"{$id}",
		start:1,
		end:9,
	}
	r3 := resource{
		url:"http://gxcms.lo/2/{$id}.html",
		target:"{$id}",
		start:1,
		end:31,
	}
	res = append(res,r1, r2, r3)
	return  res
}

func buildUrl(res []resource)[]string  {
	var list []string
	for _, v := range res {
		for i:=v.start;i<=v.end;i++ {
			if len(v.target) > 0 {
				url := strings.Replace(v.url,v.target,strconv.Itoa(i),1)
				list = append(list, url)
			} else {
				list = append(list, v.url)
			}
		}
	}
	return list
}

func makeLog(currentUrl,refer, ua string)string  {
	u := url.Values{}
	u.Set("time",time.Now().Format("2006-01-02 15:04:05"))
	u.Set("url",currentUrl)
	u.Set("refer",refer)
	u.Set("ua",ua)
	paramUrl := u.Encode()
	logTemplate := "127.0.0.1 - - [%s] \"GET /dot.gif?%s HTTP/1.1\" 200 43 \"%s\" \"%s\" \"-\"\n"
	logStr := fmt.Sprintf(logTemplate,time.Now().Format("2006-01-02 15:04:05"),paramUrl,refer,ua)
	return logStr

}

// 命令行运行 go run run.go --total=1000 --filePath=/tmp/xxx
func main() {

	// 获取命令行参数
	total := flag.Int("total", 1000,"抓取行数")
	filePath := flag.String("filePath","/data/log/tongji_access.log", "文件路径")
	flag.Parse()

	fmt.Print(*total, *filePath)

	// 构造网站url
	res := ruleResource()

	list := buildUrl(res)
	fmt.Println(list)



	i := *total
	for i > 0 {
		i--
		rand.Seed(time.Now().UnixNano())
		curl := list[rand.Intn(len(list))]
		refer := list[rand.Intn(len(list))]
		logStr := makeLog(curl,refer,"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36")
		//fmt.Print(logStr)
		err := ioutil.WriteFile(*filePath, []byte(logStr),0644)
		if err !=nil {
			fmt.Println(err)
			continue
		}
	}

	// 按照要求生成total 行日志内容
	fmt.Println("success")
}
