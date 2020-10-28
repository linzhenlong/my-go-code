package main

import (
	"fmt"
	"sync"
)

func main() {

	ch := make(chan []int, 100)
	res := &sync.Map{}
	//
	var wg sync.WaitGroup
	wg.Add(10)
	go func() {
		var i int
		//模拟MySQL取设备
		for {
			if i > 100 {
				close(ch)
				break
			}
			deviceids := []int{i}
			ch <- deviceids
			i++
		}
	}()

	workerNum := 10
	for i := 0; i < workerNum; i++ {
		// 模拟请求梁栋接口
		go func() {
			for dids := range ch {
				for _, v := range dids {
					res.Store(v, v)
				}
				//fmt.Println(dids)
			}
			wg.Done()
		}() 
	}
	wg.Wait()

	res.Range(func(k, v interface{}) bool {
		fmt.Println(v)
		return true
	})

}
