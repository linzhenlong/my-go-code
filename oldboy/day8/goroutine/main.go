package main

import(
	"fmt"
	"time"

)

func test() {
	var i int 
	for {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 100)
		i++
	}
}
func main() {
	go test()
	time.Sleep(time.Second)
}