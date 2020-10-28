package main

import "math"

import "fmt"

func max(a,b int64) int64 {
	return int64(math.Max(float64(a),float64(b)))
}

func main() {
	fmt.Println(max(1,2))
	a := math.MaxInt64-2
	b := math.MaxInt64-1
	fmt.Println("math.MaxInt64-2=",a)
	fmt.Println("math.MaxInt64-1=",b)
	fmt.Println("int64(math.MaxInt64-2)=",int64(a))
	fmt.Println("int64(math.MaxInt64-1)=",int64(b))
	fmt.Println("max(math.MaxInt64-2,math.MaxInt64-1)=",max(math.MaxInt64-2,math.MaxInt64-1))
	fmt.Println("math.Max(float64(a),float64(b)=",math.Max(float64(a),float64(b)))
	fmt.Println("int64(math.Max(float64(a),float64(b))=",int64(math.Max(float64(a),float64(b))))
}