package main

import (
	"testing"
) // 引入go的testing 框架

//编写一个测试用例去测试addUpper函数是否正确

func TestAddUpper(t *testing.T) {
	// 调用
	res := addUpper(10)
	if res != 55 {
		// 输出并退出
		t.Fatalf("addUpper()有误 返回值%d期望值%d\n", res,55)
	}
	// 如果正确输出日志
	t.Logf("addUpper()正确 返回值%d期望值%d\n", res,55)
}

func TestSum(t *testing.T)  {
	res := sum(1,9)
	if res != 10 {
		t.Fatalf("sum(1,9)有误,期望值为：%d ，实际值为：%d\n",10, res)
	}
	t.Logf("sum(1,9)正确,期望值为：%d ，实际值为：%d\n",10, res)
}