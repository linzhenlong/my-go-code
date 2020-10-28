package main

import (
	"fmt"
	"strconv"
	"sync"
)

type userAges struct {
	ages map[string]int

	// sync.Mutex // 互斥锁
	// 将上面互斥锁，改完读写锁
	sync.RWMutex
}

func (u *userAges) add(name string, age int) {
	u.Lock()
	defer u.Unlock()
	u.ages[name] = age
}
func (u *userAges) get(name string) int {

	// 这个地方，读不加锁，可能会报错
	// 当没值呢，就开始读了,可以加个读写锁
	u.RLock()
	defer u.RUnlock()

	if age, ok := u.ages[name]; ok {
		return age
	}
	return -1
}

func main() {
	userAges := userAges{
		ages: map[string]int{},
	}
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			userAges.add(strconv.Itoa(i), i)
			wg.Done()
		}(i)
	}
	
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(userAges.get(strconv.Itoa(i)))
			wg.Done()
		}(i)
	}
	// 阻塞一下主协程
	wg.Wait()
}
