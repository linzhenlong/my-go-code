package main

import "time"

import "log"

func main() {
	// 1.获取ticker对象
	ticker1 := time.NewTicker(time.Second * 1)

	// 子协程
	go func() {
		i := 0
		for {
			i++
			<-ticker1.C
			log.Println("ticker1...")
			if i > 5 {
				ticker1.Stop()
			}
		}
	}()
	for {
	}
}
