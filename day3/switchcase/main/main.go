package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	var rand_num int = rand.Intn(100)
	for {
		flag := false
		var a int
		fmt.Scanf("%d", &a)
		switch {
		case a == rand_num:
			fmt.Println("你猜对啦")
			flag = true
		case a > rand_num:
			fmt.Println("大了")
		case a < rand_num:
			fmt.Println("小了")
		}
		if flag {
			break
		}
	}
}
