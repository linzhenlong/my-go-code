package main

import "fmt"

func main() {
	sn1 := struct {
		age  int
		name string
	}{age: 10, name: "张三"}
	sn2 := struct {
		age  int
		name string
	}{age: 10, name: "张三"}

	if sn1 == sn2 {
		fmt.Println("sn1==sn2")
	}
	m := make(map[string]string)
	m["a"] = "1"
	sm1 := struct {
		age int
		m   map[string]string
	}{age: 10, m: m}

	sm2 := struct {
		age int
		m   map[string]string
	}{age: 10, m: m}

	// 那什么是可比较的呢，常见的有 bool、数值型、字符、指针、数组等，像切片、map、函数等是不能比较的。
	// 具体可以参考 Go 说明文档。https://golang.org/ref/spec#Comparison_operators

	//if sm1 == sm2 {
	fmt.Println("sm1==sm2", sm1, sm2)
	//}
}
