package main

import "context"

import "time"

import "fmt"

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("done...")
				return
			default:
				fmt.Println("睡一秒")
				time.Sleep(time.Second)
			}
		}
	}()
	time.Sleep(time.Second * 5)
	fmt.Println("mian 睡五秒...")
	cancel()

}
