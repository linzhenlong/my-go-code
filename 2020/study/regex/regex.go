package main

import(
	"regexp"
	"fmt"

)
const text = `My email is linzl@gmail.com~~~
			email1 is email@qq.com
			email2 is test@def.org.cn
			email3 is test@sohu.com`
func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)(\.[a-zA-Z0-9]+)`)
	//match := re.FindAllString(text, -1)
	match := re.FindAllStringSubmatch(text, -1)
	//fmt.Println(match)

	for _, m := range match {
		for _, val := range m {
			fmt.Printf("%20s",val)
		}
		fmt.Println()
	}
}