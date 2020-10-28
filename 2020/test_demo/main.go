package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/techoner/gophp/serialize"
)

func main() {

	str := `a:1:{i:0;a:4:{s:9:"pro_price";s:9:"￥109.53";s:8:"pro_mall";s:15:"美国亚马逊";s:7:"pro_url";s:36:"https://www.amazon.com/dp/B0013OXE46";s:7:"mall_id";i:41;}}`

	// unserialize() in php
	out, _ := serialize.UnMarshal([]byte(str))

	fmt.Println(out) //map[php:世界上最好的语言]

	// serialize() in php
	jsonbyte, _ := serialize.Marshal(out)

	fmt.Println(string(jsonbyte)) // a:1:{s:3:"php";s:24:"世界上最好的语言";}

	fmt.Println(math.Pow(2, 8))

	var funcs []func(int, int) int
	funcs = append(funcs, test1)
	funcs = append(funcs, test2)
	funcs = append(funcs, test3)
	funcs = append(funcs, test4)

	for _, f := range funcs {
		fmt.Println(f(10, 9))
	}

	str2 := "ddddddd"

	fmt.Println(strings.Split(str2, ","))

}

func test1(n, j int) int {
	fmt.Println("test1")
	return n + j
}
func test2(n, j int) int {
	fmt.Println("test2")
	return 2*n + j
}
func test3(n, j int) int {
	fmt.Println("test3")
	return n - j
}
func test4(n, j int) int {
	fmt.Println("test4")
	return 3*n + j
}
