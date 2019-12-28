package main

import (
	"fmt"
	"testing"
)
// 文件名必须以xxxx_test.go 的格式命名


// 使用TestMain作为初始化test,
// 并且使用m.Run()来调用其他tests，
// 可以完成一些需要初始化操作的testing,
// 比如数据库连接，文件打开，REST服务登录等

// 如果没有在TestMain中调用m.Run()
// 则除了TestMain以为的其他tests都不会被执行
func TestMain(m *testing.M)  {
	fmt.Println("test main first")

	m.Run()
}

// 测试用例的case 必须是以Test开头符合Testxxx的形式
func TestPrint1(t *testing.T)  {

	sum := sum(10,20)
	if sum != 30 {
		t.Fatal("计算错误")
	} else {
		testAbc(t)
	}
}

func TestPrint2(t *testing.T)  {
	// t.SkipNow()为跳过当前的test,
	// 并且直接按照PASS处理继续下一个，
	// 并且需要写在测试用例的第一行，中间或是末尾不起作用
	t.SkipNow()
	sum := sum(10,20)
	if sum != 30 {
		t.Fatal("计算错误")
	} else {
		testAbc(t)
	}
}

// 使用t.Run()来执行subtests,可以控制test的输出及test的顺序
func TestTRun(t *testing.T)  {
	t.Run("a1", func(t *testing.T) {
		fmt.Println("a1")
	})
	t.Run("a2",a2)
	t.Run("a3",a2)
}

func a2(t *testing.T)  {
	fmt.Println("a2")
}

func testAbc(t *testing.T)  {
	fmt.Print("success")
}

// benchmark


// 1. benchmark函数一般以Benchmark开头
// 2. benchmark的case一般会跑b.N次，而且每次执行都是如此
// 3. 在执行过程中会根据实际的case的执行时间是否稳定会增加b.N次数以达到稳定

// go test -bench=. -v
func BenchmarkSum2(b *testing.B) {
	for n:=0;n<b.N;n++ {
		sum2(10)
	}
}
