package main

import "fmt"

import "bufio"

import "os"

import "strings"

func bufioDemo() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入内容:")
	text , _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	fmt.Printf("%#v\n", text)
}

func main() {
	var (
		name string
		age int
		married bool
	)
	/* _, err := fmt.Scan(&name, &age, &married)
	if err != nil {
		panic("scan error")
	}
	fmt.Printf("扫描结果 name:%s,age:%d,married:%t\n", name, age, married) */

	//fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
	fmt.Scanln(&name, &age, &married)
	fmt.Printf("扫描结果:name:%s age:%d married:%t \n", name, age, married)
	bufioDemo()
}