package main

import (
	"fmt"
	"time"
)

func test(i int)  {
	fmt.Println(i)
}

func main()  {

	for i:=0;i<=100;i++  {
		go test(i)
	}

	time.Sleep(time.Second * 10)
}
