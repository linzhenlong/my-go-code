package main

import (
	"testing"
)

// go test -coverprofile=c.out 跑代码覆盖率
// go tool cover -html=c.out  覆盖率文件.
func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		{"abcabcbb", 3},
		{"pwwkew", 3},
		{"", 0},
		{"b", 1},
		{"bbbbb", 1},
		{"bbbbbcb", 2},
		{"abcabcabcd", 4},
		{"这里是中国北京", 7},
		{"一二三二一", 3},
		{"黑灰化肥灰会挥发发灰黑讳为黑灰花会飞", 7},
	}
	for _, tt := range tests {
		actual := lengthOfNoRepeatingSubStr(tt.s)
		if actual != tt.ans {
			t.Errorf("got %d for input %s expected %d", actual, tt.s, tt.ans)
		}
	}
}

// 1.跑用例:go test -bench .
// 2.go test -bench . -cpuprofile cpu.out 看cpu使用率

// 3. go tool pprof cpu.out 查看cpu.out.
// 3.1 输入web 可以看到
func BenchmarkSubstr(b *testing.B) {
	s := "黑灰化肥灰会挥发发灰黑讳为黑灰花会飞"
	// 把s扩大几倍
	for i:=0; i< 10;i++ {
		s +=s
	}
	b.Logf("len(s)=%d", len(s))
	b.ResetTimer() // 重置一下时间，将准备数据的时间排除出去.
	ans := 7
	for i := 0; i < b.N; i++ {
		actual := lengthOfNoRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("got %d for input %s expected %d", actual, s, ans)
		}
	}
}
// 8-3 