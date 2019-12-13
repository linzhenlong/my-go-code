package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Rect struct {
	X int
	Y int
}

type Integer int

/**
画一个10x8的矩形
 */
func (r *Rect) print() {
	(*r).X = 10
	(*r).Y = 8

	for i:=0;i<(*r).X;i++{
		for j:=0;j<(*r).Y;j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}


func (num *Integer) isOddNumber() {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(1000)
	if i%2 == 0 {
		fmt.Printf("i=%d是偶数\n",i)
	} else {
		fmt.Printf("i=%d是奇数\n",i)
	}
}

/**
	根据行,列，字符打印,对应的行数，列数的字符
	比如：行:3,列:2,字符*
 */
func (num *Integer) Print2(n int,m int,key string) {
	for i:=1;i<=n;i++  {
		for j:=1;j<m;j++ {
			fmt.Print(key)
		}
		fmt.Println()
	}
}

// 画一个 nxm的矩形
func (r *Rect) print2(n int, m int)  {
	(*r).X = n
	(*r).Y = m

	for i:=0;i<(*r).X;i++{
		for j:=0;j<(*r).Y;j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (num *Integer) Jisuanqi(n float64, m float64,action string) (res float64) {
	switch action {
	case "+":
		res = n + m
	case "-":
		res = n - m
	case "*":
		res = m * n
	case "/":
		res = n/m
	default:
		res = 0
	}
	return
}

// 求矩形的面积
func (r *Rect) area(len int, width int) int {
	(*r).X = len
	(*r).Y = width
	return (*r).X * (*r).Y
}

func main()  {

	var r Rect
	(&r).print()

	fmt.Println("=============")

	var r2 Rect
	(&r2).print2(4,5)

	area :=(&r2).area(4,5)
	fmt.Println("面积:",area)

	var i Integer
	(&i).isOddNumber()

	i.Print2(5,6,"%%%%")

	res := i.Jisuanqi(10.0,1.0,"/")
	fmt.Println(res)

	/**
		将二维数组[[1 2 3] [4 5 6] [7 8 9]] 转化成[[1 4 7] [2 5 8] [3 6 9]]
	 */
	var arr  = [3][3]int{{1,2,3},{4,5,6},{7,8,9}}
	var arr2 [3][3]int
	for i ,_ := range arr {
		for j,v2 := range arr[i] {
			fmt.Printf("arr[%d][%d] = %d \t",i,j,v2,)
			arr2[j][i] = v2
		}
		fmt.Println()
	}
	fmt.Println(arr)
	fmt.Println(arr2)
}


