package main
import (
	"errors"
	"math"
	"runtime"
	"reflect"
	"fmt"
)
func eval(a,b int,op string) (ressult int, err error) {
	switch op {
	case "+":
		ressult = a + b
	case "-":
		ressult = a - b 
	case "*":
		ressult = a * b
	case "/":
		ressult = a / b
	default:
		//panic("unsupported operator:"+op) // panic 会让程序报错，并退出来
		err = errors.New("unsupported operator:"+op)
	}
	return
}
// 带余数除法，13/3 = 4余1
func div(a,b int)(int, int) {
	return a/b, a%b
}
// 重写pow 函数
func pow(a,b int)int {
	return int(math.Pow(float64(a), float64(b)))
}
func apply(op func(int, int)int,a,b int) int{
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("calling function %s with args (%d,%d)\n",opName,a,b)
	return op(a,b)
}

// 函数的可变参数列表
func sum(numbers ...int) int {
	s := 0;
	for i := range numbers {
		s += numbers[i]
	}
	return s
}
// 指针相关的
// 交换两个变量的值
func swap(a, b int) {
	b, a = a, b
}

func swap2(a, b *int) {
	*b, *a = *a, *b
}

func swap3(a, b int) (int, int) {
	return b, a
}
func main() {
	fmt.Println(eval(1, 1, "+"))
	fmt.Println(div(13,3))
	fmt.Println(apply(pow,3,4))

	fmt.Println(apply(func(a,b int) int{
		return a *b
	},3,4))
	fmt.Println(sum(1,2,3,4))

	a, b := 3, 4
	swap(a, b)
	fmt.Println(a, b) // 3 4
	swap2(&a, &b)
	fmt.Println(a, b) // 4 3
	swap3(a, b)
	fmt.Println(a,b)
	fmt.Println(swap3(a,b))
}