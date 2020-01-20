package main

import "fmt"

import "math/cmplx"

import "math"

// 变量默认值
func variableZeroValue() {
	var a int
	var b string
	fmt.Printf("%d,%s,%q\n", a, b, b) // 0,,""
}
// 变量初始值
func variableInitialValue() {
	var a ,b int = 3, 10
	var s string = "hello"
	fmt.Println(a, s, b)
}
// 变量类型推导.
func variableTypeDeduction() {
	var a,b,c,d = 2,3,"我是字符串",true
	fmt.Println(a,b,c,d)
}


// 省略var 用:=定义变量
func variableShorter() {
	 a,b,c,d := 2,3,"我是字符串",true
	fmt.Println(a,b,c,d)
}
// 函数内定义的变量，作用域只在函数体内部，叫做局部变量

// 函数外定义的变量叫做包内部变量,作用是包内部
// 函数外部不能用 :=定义变量
var (
	aa = "我是全局变量"
	kk = 666
)

func variablePackageGlobal() {
	fmt.Printf("函数内使用包内全局变量,aa=%s,kk=%d\n", aa, kk)
	aa = "我在给你改了，看下作用域"
	kk = 8888
}

// 内建变量类型
// 1. bool,string 
// 2. (u)int,(u)int8, (u)int32,(u)int64 uintprt(指针类型)
// 3. byte 字节型，run 字符型，相当于char
// 浮点数:float32,float64,complex64(复数),complex128(复数)

// 复数 :i = 根号下-1
// 复数 : 3(实部) + 4i(虚部)

func eular() {
	// 复数
	/* c := 3 + 4i
	fmt.Printf("复数:3 + 4i =%v\n",cmplx.Abs(c)) */
	eular1 := cmplx.Pow(math.E, 1i * math.Pi) + 1 //e的一π次方+1
	eular2 := cmplx.Exp(1i * math.Pi )+1 //e的一π次方+1
	fmt.Println("欧拉公式:",eular1)
	fmt.Printf("欧拉公式:%.3f\n",eular1)
	fmt.Println("欧拉公式:",eular2)
	fmt.Printf("欧拉公式:%.3f\n",eular2)
}

// 强制类型转换
// 1. 类型转换是强制的
// 2. var a,b int = 3,4 ,把a,b作为直角三角形的两个边，求斜边
// 
func triangle() {
	var a, b int = 3,4
	var c int
	c = int(math.Sqrt(float64(a * a + b * b)))
	fmt.Println(c)
}

// 常量的定义
func consts() {
	const filename string = "abc.txt"
	const a, b = 3,4
	var c int 
	c = int(math.Sqrt(a * a + b * b))
	fmt.Println(filename, c)
}

// 枚举类型常量
func enums() {

	// 普通枚举类型
	/* const(
		cpp = 0
		java = 1
		python = 2
		golang =3
	) */

	// 自增枚举类型
	const (
		cpp = iota //自增值
		java
		_  // 跳过一个
		python
		golang
	)
	fmt.Println(cpp,java,python,golang)
	// 定义b,kb,mb,gb,tb,pb...
	const (
		b = 1 << (10 * iota) // 左移10*iota位
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
	
}

func main() {
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	fmt.Printf("全局变量==> aa=%s,kk=%d\n",aa, kk)
	variablePackageGlobal()
	fmt.Printf("全局变量==> aa=%s,kk=%d\n",aa, kk)

	// 欧拉公式
	eular()
	triangle()
	consts()
	enums()
}
//该看/Users/smzdm/Documents/自己/狗浪/【提升高度】Google资深工程师带你全面掌握Go语言/第2章 基础语法/2-2 内建变量类型.avi