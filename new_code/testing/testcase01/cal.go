package main

// 被测试的函数
func addUpper(n int) int {
	res := 1
	for i:=1;i<=n;i++ {
		res +=i
	}
	return res
}

func sum(a int,b int) int  {
	return  a + b
}

//该看视频260了