package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"
)

var (
	// 图片正则表达式
	reImage = `< *[img][^\\>]*[src] *= *[\\"\\']{0,1}([^\\"\\'\\ >]*)`

	// 数据管道
	chanImageURLs chan string
	wg            sync.WaitGroup

	// 监控协程
	chanTask chan string
)

// GetPage ...
func GetPage(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	pageBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(pageBytes), nil
}
// GetImg 获取图片... 
func GetImg(url string) {
	htmlBodyStr, err := GetPage(url)
	HandleError(err, "GetPage")
	// 过滤数据
	//regexp :=  regexp.MustCompile(reQQEmail)
	regexp := regexp.MustCompile(reImage)
	res := regexp.FindAllStringSubmatch(htmlBodyStr, -1)

	for _, v := range res {
		chanImageURLs <- v[1]
	}
	chanTask <- url
	wg.Done()
}

// DownloadFile 下载图片...
func DownloadFile(url string, fileName string) (ok bool) {
	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	imageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "DownloadFile ioutil.ReadAll")
	fileName = "./image/" + fileName

	err = ioutil.WriteFile(fileName, imageBytes, 0666)
	if err != nil {
		fmt.Println(err)
		ok = false
	}
	ok = true
	return true
}

// HandleError 处理异常
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}
// TaskCheck 检测任务
func TaskCheck() {
	var count int
	for {
		url := <-chanTask
		fmt.Printf("%s 完成爬取任务\n", url)
		count++
		if count == 26 {
			close(chanImageURLs)
			break
		}
	}
	wg.Done()
}

// 并发爬图片思路
// 1.初始化一个数据管道
// 2.爬虫协程26个(https://www.umei.cc/p/gaoqing/cn/1.htm 有26页数据),26个协程往管道添加图片链接.
// 3.任务统计协程:检查26个任务是否都完成，完成则关闭数据管道
// 4.下载协程:从管道读取链接并下载.
func main() {
	//GetImg("https://www.umei.cc/p/gaoqing/cn/1.htm")
	//DownloadFile("http://i1.shaodiyejin.com/uploads/tu/201908/10320/349df80c83_33.jpg","1.jpg")

	// 1.初始化管道
	chanImageURLs = make(chan string, 100)
	chanTask = make(chan string, 26)

	// 2.爬虫协程
	for i := 1; i < 27; i++ {
		wg.Add(1)
		url := fmt.Sprintf("https://www.umei.cc/p/gaoqing/cn/%d.htm", i)
		go GetImg(url)
	}

	// 2. 任务统计协程
	wg.Add(1)
	go TaskCheck()

	// 3.下载协程，从数据管道读取并下载
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for url := range chanImageURLs {
				//nano := time.Now().UnixNano()
				path := time.Now().Format("15-04-05")
				fileName := fmt.Sprintf("%s_%s.jpg", path, getFileNameFromURL(url))
				DownloadFile(url, fileName)
			}
		}()
	}
	wg.Wait()
	fmt.Println("main over")
}

func getFileNameFromURL(url string) string {
	strs := strings.Split(url, "/")
	return strs[len(strs)-1]
}
