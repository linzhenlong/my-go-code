package main

import "encoding/base64"

import "fmt"

import "hash/crc32"

import "sort"

import "time"

import "log"

// ServerSlice ... 实现sort接口
type ServerSlice []Server

// Server ...
type Server struct {
	Weight int
}

func (p ServerSlice) Len() int {
	return len(p)
}
func (p ServerSlice) Less(i, j int) bool {
	return p[i].Weight > p[j].Weight
}
func (p ServerSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Sort is a convenience method.

func main() {
	str := "linzl:123"
	base64Str := base64.StdEncoding.EncodeToString([]byte(str))
	fmt.Println(base64Str)

	// ip_hash 算法
	ip := "localhost:9092"
	fmt.Println(crc32.ChecksumIEEE([]byte(ip)))

	nums := []int{5, 6, 8, 9, 1, 2, 3}
	sort.Ints(nums)
	fmt.Println(nums)

	ss := ServerSlice{
		Server{Weight: 4},
		Server{Weight: 1},
		Server{Weight: 2},
	}
	sort.Sort(ss)
	fmt.Println(ss)

	// 定时器.
	t := time.NewTicker(time.Second * 2)
	for {
		select {
		case <-t.C:
			log.Println("aaaa")
		}
	}
}
