package main

import "testing"

// benchmark


// 1. benchmark函数一般以Benchmark开头
// 2. benchmark的case一般会跑b.N次，而且每次执行都是如此
// 3. 在执行过程中会根据实际的case的执行时间是否稳定会增加b.N次数以达到稳定

// go test -bench=. -v
func BenchmarkAll(b *testing.B) {
	for n:=0;n<b.N;n++ {
		sum(10)
	}
}
