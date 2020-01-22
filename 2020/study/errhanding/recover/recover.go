package main

import (
	"fmt"
)
// 仅在defer调用中使用
// 获取panic的值
// 如果无法处理。可以重新panic
func tryRecover() {
	defer func() {
		r := recover() //返回interface{}
		// 类型断言.
		if err , ok := r.(error);ok {
			fmt.Println("error occurred:", err)
		} else {
			// 无法处理重新panic
			panic(fmt.Sprintf("不知道什么错误:%v", r))
		}
	}()
	//panic(errors.New("this is an error"))
	/* b := 0
	a := 5/b // 故意写错了
	fmt.Println(a) */
	panic(1223)
}

func main() {
	tryRecover()
}
