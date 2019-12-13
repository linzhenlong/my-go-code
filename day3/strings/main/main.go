package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main()  {
	var url string = "baidu.com/baidu.com/baidu.com"

	// 判断url是否以"http://"开头
	if !strings.HasPrefix(url,"http://") {
		url = "http://"+url
	}
	// 判断url是否以"/"结尾
	if !strings.HasSuffix(url,"/") {
		url = url + "/"
	}
	// 返回baidus 在url首次出现的位置,如果url不包含baidus 则返回-1
	indexNum := strings.Index(url,"baidus")
	fmt.Printf("strings.Index(url,\"/baidus\")的最后位置:%d\n",indexNum)

	// 返回"/"在url最后出现的位置,如果url不包含"/",则返回-1
	lastNum := strings.LastIndex(url,"/")

	fmt.Printf("strings.LastIndex(url,\"/\")的最后位置:%d\n",lastNum)
	fmt.Printf("url is %s\n",url)

	// baidu在http://baidu.com/baidu.com/baidu.com/中一共出现strings.Count(url,"baidu")次
	baiduCount := strings.Count(url,"baidu")
	fmt.Printf("baidu在%s中一共出现%d次\n",url,baiduCount)

	// 将url中的baidu替换成g,替换2个,替换后url:http://g.com/g.com/baidu.com/
	url = strings.Replace(url,"baidu","g",2)
	fmt.Printf("将url中的baidu替换成g,替换2个,替换后url:%s\n",url)

	// str:banana,strings.Repeat("na",2) 后str:banana
	var str string = "ba"
	str = str + strings.Repeat("na",2)
	fmt.Printf("str:%s,strings.Repeat(\"na\",2) 后str:%s\n",str,str)

	// str:banana,strings.ToUpper(str) 转大写后ToUpperStr:BANANA
	ToUpperStr := strings.ToUpper(str);
	fmt.Printf("str:%s,strings.ToUpper(str) 转大写后ToUpperStr:%s\n",str,ToUpperStr)

	// ToUpperStr:BANANA,strings.ToLower(ToUpperStr)转小写后ToUpperStr:banana
	fmt.Printf("ToUpperStr:%s,strings.ToLower(ToUpperStr)转小写后ToUpperStr:%s\n",ToUpperStr,strings.ToLower(ToUpperStr))

	// 通过strings.TrimSpace 去掉" 您 好，傻x " 的首尾空格:您 好，傻x
	fmt.Printf("通过strings.TrimSpace 去掉\" 您 好，傻x \" 的首尾空格:%s\n",strings.TrimSpace(" 您 好，傻x "))

	// 通过strings.Trim 去掉" 您 好，傻x " 中的"傻x":您 好，
	fmt.Printf("通过strings.Trim 去掉\" 您 好，傻x \" 中的\"傻x\":%s\n",strings.Trim(" 您 好，傻x ","傻x "))

	// 通过strings.TrimLeft 去掉" 您 好，傻x "中左边的空格:您 好，傻x
	fmt.Printf("通过strings.TrimLeft 去掉\" 您 好，傻x \"中左边的空格:%s\n",strings.TrimLeft(" 您 好，傻x "," "))

	// 通过strings.TrimRight 去掉" 您 好，傻x!"中右边的x!: 您 好，傻
	fmt.Printf("通过strings.TrimRight 去掉\" 您 好，傻x!\"中右边的x!:%s\n",strings.TrimRight(" 您 好，傻x!","x!"))

	// 字符"a b c d e"通过strings.Fields转数组:["a" "b" "c" "d" "e"] ,strings.Fields返回空格分隔的子串slice
	fmt.Printf("字符\"a b c d e\"通过strings.Fields转数组:%q\n",strings.Fields("a b c d e"))

	// 字符串"A,B,c;d"通过strings.Split("A,B,c;d",",")切分为:["A" "B" "c;d"]
	fmt.Printf("字符串\"A,B,c;d\"通过strings.Split(\"A,B,c;d\",\",\")切分为:%q\n",strings.Split("A,B,c;d",","))

	s := []string{"abc","efg","h","i","jk"}
	// ["abc" "efg" "h" "i" "jk"]通过strings.Join(s,",") 转成字符串:abc,efg,h,i,jk
	fmt.Printf("%q通过strings.Join(s,\",\") 转成字符串:%s\n",s,strings.Join(s,","))

	// 整形装字符串
	number := 1024
	fmt.Println(strconv.Itoa(number))

	// 字符串转整数
	str_num := "1204"
	fmt.Println(strconv.Atoi(str_num))
}