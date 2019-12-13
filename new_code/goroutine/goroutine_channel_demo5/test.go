package main

import (
	"fmt"
	"time"
)

func main()  {


	var slice []int
	start := time.Now().Unix()
	for i:=1;i<=80000;i++ {
		flag := true
		for j:=2;j<=i;j++ {
			if i%j == 0 {
				flag = false
				break
			}
		}
		if flag {
			slice = append(slice, 1)
		}
	}
	end := time.Now().Unix()
	fmt.Println("run_time:",end-start)
}
