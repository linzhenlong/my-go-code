package main

import (
	"fmt"
	"regexp"
	"io/ioutil"
	"net/http"
)
var (
	reQQEmail = `\d+@qq.com`
	// \w 代表大小写字母+数字+下划线
	reEmail = `\w+@\w+.\w+`
	emailURL = "https://tieba.baidu.com/p/6051076813?red_tag=2583501189"
	reLink = `href="(.*?)"` //
	// +代表出一次或是多次
	// \s\S 代表各种字符
	// +? 代表贪婪模式,
	reLink2 = `href="(https?://[\s\S]+?)"` //
	
	//rePhone = `1[3456789]\d\s?\d{4}\s?\d{4}`
	rePhone = `<span>(1[3456789]\d)</span><span>(\d{4})</span><span>(\d{4})</span>`

	// 
	reIDCard = `[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]`

	reImage = `< *[img][^\\>]*[src] *= *[\\"\\']{0,1}([^\\"\\'\\ >]*)`

)

// GetEmail 爬邮箱
func GetEmail(url string) {
	// resp, err := http.Get("https://tieba.baidu.com/p/6051076813?red_tag=2583501189")
	// HandleError(err, "http.Get")
	// defer resp.Body.Close()
	// htmlBodyBytes,err := ioutil.ReadAll(resp.Body)
	// HandleError(err, "ioutil.ReadAll")
	htmlBodyStr,err := GetPage(url)
	HandleError(err, "GetPage")
	// 过滤数据
	//regexp :=  regexp.MustCompile(reQQEmail)
	regexp :=  regexp.MustCompile(reEmail)
	res := regexp.FindAllStringSubmatch(htmlBodyStr,-1)
	var emailList []string
	for _,v := range res {
		for _, v2 := range v {
			emailList = append(emailList, v2)
		}
	}
	fmt.Println(emailList)

}

// GetPage ...
func GetPage(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	pageBytes, err := ioutil.ReadAll(resp.Body)
	if err !=nil {
		return "", err
	}
	return string(pageBytes),nil
}
// GetLink 获取页面url...
func GetLink(url string) {
	htmlBodyStr,err := GetPage(url)
	HandleError(err, "GetPage")
	// 过滤数据
	//regexp :=  regexp.MustCompile(reQQEmail)
	regexp :=  regexp.MustCompile(reLink2)
	res := regexp.FindAllStringSubmatch(htmlBodyStr,-1)
	
	for _,v := range res {
		fmt.Println(v[1])
	}
}
// 
// GetPhone https://www.zhaohaowang.com/ ...
func GetPhone(url string) {
	htmlBodyStr,err := GetPage(url)
	HandleError(err, "GetPage")
	// 过滤数据
	//regexp :=  regexp.MustCompile(reQQEmail)
	regexp :=  regexp.MustCompile(rePhone)
	res := regexp.FindAllStringSubmatch(htmlBodyStr,-1)
	
	for _,v := range res {
		fmt.Println(v)
	}
}
func GetIDCard(url string) {
	htmlBodyStr,err := GetPage(url)
	HandleError(err, "GetPage")
	// 过滤数据
	//regexp :=  regexp.MustCompile(reQQEmail)
	regexp :=  regexp.MustCompile(reIDCard)
	res := regexp.FindAllStringSubmatch(htmlBodyStr,-1)
	
	for _,v := range res {
		fmt.Println(v)
	}
}

func GetImg(url string) {
	htmlBodyStr,err := GetPage(url)
	HandleError(err, "GetPage")
	// 过滤数据
	//regexp :=  regexp.MustCompile(reQQEmail)
	regexp :=  regexp.MustCompile(reImage)
	res := regexp.FindAllStringSubmatch(htmlBodyStr,-1)
	
	for _,v := range res {
		fmt.Println(v)
	}
}
//https://www.umei.cc/p/gaoqing/cn/1.htm

// HandleError 处理异常
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}
func main() {
	//GetEmail(emailURL)
	//GetLink("https://www.smzdm.com")
	//GetPhone("https://www.zhaohaowang.com/")
	//GetIDCard("http://www.yangtse.com/content/776869.html")
	GetImg("https://www.umei.cc/p/gaoqing/cn/1.htm")
}
