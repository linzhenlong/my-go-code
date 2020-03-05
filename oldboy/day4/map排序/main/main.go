package main

import (
	"fmt"
	"sort"
)

func main() {
	testMapSort()
	test2()
}

func testMapSort() {
	var a = make(map[int]int, 5)
	a[8] = 10
	a[3] = 10
	a[2] = 10
	a[1] = 10
	a[18] = 10

	var keys []int
	for key, _ := range a {
		keys = append(keys, key)
	}
	fmt.Println(keys)
	sort.Ints(keys)
	fmt.Println(keys)
	for _, v := range keys {
		fmt.Println(v, a[v])
	}
}

func test2() {
	var a [26]byte
	a[0] = 'A'
	for i := 1; i < 26; i++ {
		a[i] = byte(a[i-1] + 1)
	}
	fmt.Printf("%q", a)
}
