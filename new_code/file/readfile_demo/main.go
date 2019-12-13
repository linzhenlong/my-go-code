package main

import (
	"fmt"
	"io/ioutil"
)

const READ_FILE_NAME = "/Users/smzdm/Desktop/10.255.19.90.txt"
func main()  {

	// 使用ioutil.ReadFile 一次性读取文件,适用于小文件
	content, err := ioutil.ReadFile(READ_FILE_NAME)
	if err !=nil {
		fmt.Println(err)
	}
	fmt.Printf("%v", content)
	fmt.Println()
	fmt.Printf("%s", content)

}
