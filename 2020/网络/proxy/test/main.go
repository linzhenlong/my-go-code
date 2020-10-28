package main

import "encoding/base64"

import "fmt"

func main() {
	str := "linzl:123"
	base64Str := base64.StdEncoding.EncodeToString([]byte(str))
	fmt.Println(base64Str)

}