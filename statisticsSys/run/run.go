package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type resource struct {
	url    string
	target string
	start  int
	end    int
}

var ualist = []string{
	"Mozilla/5.0 (iPod; U; CPU iPhone OS 4_3_2 like Mac OS X; zh-cn) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8H7 Safari/6533.18.5",
	"Mozilla/5.0 (iPhone; U; CPU iPhone OS 4_3_2 like Mac OS X; zh-cn) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8H7 Safari/6533.18.5",
	"MQQBrowser/25 (Linux; U; 2.3.3; zh-cn; HTC Desire S Build/GRI40;480*800)",
	"Mozilla/5.0 (Linux; U; Android 2.3.3; zh-cn; HTC_DesireS_S510e Build/GRI40) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
	"Mozilla/5.0 (SymbianOS/9.3; U; Series60/3.2 NokiaE75-1 /110.48.125 Profile/MIDP-2.1 Configuration/CLDC-1.1 ) AppleWebKit/413 (KHTML, like Gecko) Safari/413",
	"Mozilla/5.0 (iPad; U; CPU OS 4_3_3 like Mac OS X; zh-cn) AppleWebKit/533.17.9 (KHTML, like Gecko) Mobile/8J2",
	"Mozilla/5.0 (Windows NT 5.2) AppleWebKit/534.30 (KHTML, like Gecko) Chrome/12.0.742.122 Safari/534.30",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_2) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/14.0.835.202 Safari/535.1",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_2) AppleWebKit/534.51.22 (KHTML, like Gecko) Version/5.1.1 Safari/534.51.22",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 5_0 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A5313e Safari/7534.48.3",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 5_0 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A5313e Safari/7534.48.3",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 5_0 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A5313e Safari/7534.48.3",
	"Mozilla/5.0 (Windows NT 6.1) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/14.0.835.202 Safari/535.1",
	"Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; SAMSUNG; OMNIA7)",
	"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0; XBLWP7; ZuneWP7)",
	"Mozilla/5.0 (Windows NT 5.2) AppleWebKit/534.30 (KHTML, like Gecko) Chrome/12.0.742.122 Safari/534.30",
	"Mozilla/5.0 (Windows NT 5.1; rv:5.0) Gecko/20100101 Firefox/5.0",
	"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.2; Trident/4.0; .NET CLR 1.1.4322; .NET CLR 2.0.50727; .NET4.0E; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729; .NET4.0C)",
	"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.2; .NET CLR 1.1.4322; .NET CLR 2.0.50727; .NET4.0E; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729; .NET4.0C)",
	"Mozilla/4.0 (compatible; MSIE 60; Windows NT 5.1; SV1; .NET CLR 2.0.50727)",
	"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; .NET4.0E)",
	"Opera/9.80 (Windows NT 5.1; U; zh-cn) Presto/2.9.168 Version/11.50",
	"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1)",
	"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; .NET CLR 2.0.50727; .NET CLR 3.0.04506.648; .NET CLR 3.5.21022; .NET4.0E; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729; .NET4.0C)",
	"Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN) AppleWebKit/533.21.1 (KHTML, like Gecko) Version/5.0.5 Safari/533.21.1",
	"Mozilla/5.0 (Windows; U; Windows NT 5.1; ) AppleWebKit/534.12 (KHTML, like Gecko) Maxthon/3.0 Safari/534.12",
	"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; .NET CLR 2.0.50727; TheWorld)",
}

func ruleResource() []resource {
	var res []resource
	r1 := resource{
		url:    "http://gxcms.lo/",
		target: "",
		start:  0,
		end:    0,
	}
	r2 := resource{
		url:    "http://gxcms.lo/{$id}/",
		target: "{$id}",
		start:  1,
		end:    9,
	}
	r3 := resource{
		url:    "http://gxcms.lo/2/{$id}.html",
		target: "{$id}",
		start:  1,
		end:    31,
	}
	res = append(res, r1, r2, r3)
	return res
}

func buildUrl(res []resource) []string {
	var list []string
	for _, v := range res {
		for i := v.start; i <= v.end; i++ {
			if len(v.target) > 0 {
				url := strings.Replace(v.url, v.target, strconv.Itoa(i), 1)
				list = append(list, url)
			} else {
				list = append(list, v.url)
			}
		}
	}
	return list
}

func makeLog(currentUrl, refer, ua string) string {
	u := url.Values{}
	u.Set("time", time.Now().Format("2006-01-02 15:04:05"))
	u.Set("url", currentUrl)
	u.Set("refer", refer)
	u.Set("ua", ua)
	paramUrl := u.Encode()
	logTemplate := "127.0.0.1 - - [%s] \"GET /dot.gif?%s HTTP/1.1\" 200 43 \"%s\" \"%s\" \"-\""
	logStr := fmt.Sprintf(logTemplate, time.Now().Format("2006-01-02 15:04:05"), paramUrl, refer, ua)
	logStr += "\n"
	return logStr

}

func WriteLog(list []string, num int, file string, runChannel chan int) {
	var logStr string
	for i := 0; i < num; i++ {
		rand.Seed(time.Now().UnixNano())
		curl := list[rand.Intn(len(list))]
		refer := list[rand.Intn(len(list))]
		ua := ualist[rand.Intn(len(ualist))]
		logStr += makeLog(curl, refer, ua)
		//fmt.Print(logStr)
	}
	fileHandle, err := os.OpenFile(file, os.O_APPEND|os.O_RDWR, 0655)
	if err != nil {
		fmt.Println("os.OpenFile error=", err)
		return
	}
	_, _ = fileHandle.Write([]byte(logStr))
	defer fileHandle.Close()
	runChannel <- 1
}

// 命令行运行 go run run.go --total=1000 --filePath=/tmp/xxx
func main() {

	// 获取命令行参数
	total := flag.Int("total", 50000, "抓取行数")
	filePath := flag.String("filePath", "/data/log/tongji_access.log", "文件路径")
	routineNum := flag.Int("routineNum", 10, "routineNum")
	flag.Parse()

	runChannel := make(chan int, *routineNum)

	// 构造网站url
	res := ruleResource()

	list := buildUrl(res)
	//fmt.Println(list)

	st := time.Now().Unix()
	for i := 0; i < *routineNum; i++ {
		go WriteLog(list, *total / *routineNum, *filePath, runChannel)
	}

	for i:=0;i<*routineNum;i++ {
		<-runChannel
	}

	fmt.Println("over")
	ed := time.Now().Unix()
	fmt.Println("runtime:", ed-st)

}
