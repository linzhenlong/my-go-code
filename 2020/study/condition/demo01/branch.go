package main

import "io/ioutil"

import "fmt"

import "errors"

import "strconv"



// if 语句的用法
func branchOfif() {
	const filename = "./abc.txt"
	/* contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("readfile failed error:%v",err)
	} else {
		fmt.Printf("%s\n", contents)
	} */

	// if 另一种写法
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Printf("readfile failed error:%v",err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	// fmt.Printf("%s\n", contents) // 会报错，因为变量的作用域，生存期，只是在条件语句块中，类似的还有for语句
}

// switch 
// go 中switch 不用break 除非fallthrough
func eval(a, b int, op string) (ressult int, err error) {
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
func grade(score int) (res string) {
	switch {
	case score < 0 || score > 100:
		panic("Wrong score:"+strconv.Itoa(score))
	case score < 60:
		res = "F"
	case score < 80:
		res = "C"
	case score < 90:
		res = "B"
	case score <= 100:
		res = "A"
	}
	return
}

func main() {
	// if 
	branchOfif()
	// switch
	if res, err := eval(10, 20, "+");err != nil {
		fmt.Printf("eval error:%s\n", err.Error())
	} else {
		fmt.Println(res)
	}
	fmt.Println(grade(0),grade(55),grade(88),grade(95),grade(66))
}

///Users/smzdm/Documents/自己/狗浪/【提升高度】Google资深工程师带你全面掌握Go语言/第2章 基础语法/2-5 循环.avi