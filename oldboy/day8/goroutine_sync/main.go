package main

import(
	"fmt"
	"sort"
	"sync"
	_"time"

)
var ( 
	m sync.Map
	wg sync.WaitGroup
	//lock sync.Mutex
)

type task struct {
	n int
}
func calc(t *task)  {
	res := 1;
	for i:=1;i<=t.n;i++ {
		res *= i
	}
	m.Store(t.n,res)
	wg.Done()
}
func main() {
	wg.Add(100)
	for i:=1;i<=100;i++ {
		t := &task{
			n: i,
		}
		go calc(t)
	}
	//fmt.Println(m)
	//wg.Done()
	//time.Sleep(time.Second * 4)
	wg.Wait()

	var res []int
	m.Range(
		func(key, value interface{}) bool {
			res = append(res, key.(int))
			return true
		},
	)
	sort.Ints(res)
	fmt.Println(res)

	for _, value := range res {
		r,ok := m.Load(value)
		if !ok {
			continue
		}
		fmt.Printf("%d!=%d\n", value, r)
	}
}