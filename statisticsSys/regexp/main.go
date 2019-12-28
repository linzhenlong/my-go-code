package main

import (
	"fmt"
	"regexp"
)

func main() {

	str := "http://gxcms.lo/2111/"
	regexpRule := regexp.MustCompile(`/([\d]+)/(([\d]+)|([\d]?))`)

	res := regexpRule.FindStringSubmatch(str)
	fmt.Println(res)
	fmt.Println(len(res))
	fmt.Println(len(res[3]))
}
