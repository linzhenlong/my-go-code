package main

import "fmt"

func main() {
	var (
		name string
		age int
		married bool
	)
	_, err := fmt.Scan(&name, &age, &married)
	if err != nil {
		panic("scan error")
	}
	fmt.Printf("扫描结果 name:%s,age:%d,married:%t\n", name, age, married)
}