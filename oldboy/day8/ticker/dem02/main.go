package main

import (
	"log"
	"time"
)

func main() {

	//ticker := time.NewTicker(time.Second)
	tick := time.Tick(time.Second)
	ticker := time.NewTicker(time.Second * 2)

	after := time.After(time.Second * 10)

	for {
		select {
		case <-tick:
			log.Println("嘻嘻")
		case <-ticker.C:
			log.Println("哈哈")
		case <-after:
			ticker.Stop() // 用完之后需要关闭.
			log.Println("timeout")
			return
		}
	}
}
