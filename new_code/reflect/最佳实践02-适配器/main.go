package main

import (
	"fmt"
	"reflect"
)

func Test1(i,j int) int {
	return i + j
}

func Test2(i,j int,o string) float64 {
	switch o {
	case "*":
		return float64(i * j)
	case "/":
		return float64(i/j)
	case "+":
		return float64(i+j)
	case "-":
		return float64(i-j)
	default:
		return float64(i+j)
	}
}

func Bridge(f interface{}, args... interface{}) interface{}{
	var function reflect.Value
	var params []reflect.Value
	var n int
	n = len(args)
	params = make([]reflect.Value, n)
	for i:=0;i<n;i++ {
		params[i] = reflect.ValueOf(args[i])
	}
	function = reflect.ValueOf(f)
	res := function.Call(params)
	ri := res[0].Interface();
	switch ri.(type) {
	case int:
		return ri.(int)
	case float64:
		return ri.(float64)
	default:
		return 0
	}
}

func Print(n interface{})  {
	switch n.(type) {
	case int:
		fmt.Println(n.(int))
	case float64:
		fmt.Println(n.(float64))
	default:
		fmt.Println("error")
	}
}

func main()  {
	num := Bridge(Test1,8,9)
	num2 := Bridge(Test2,11,2,"*")
	Print(num)
	Print(num2)

}
