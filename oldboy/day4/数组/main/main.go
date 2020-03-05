package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	rand.Seed(time.Now().UnixNano())
	var a [10]int
	for i:=0;i<len(a);i++ {
		a[i] = rand.Intn(1000)
	}

	for key,val := range a {
		fmt.Printf("a[%d]=%d\n",key,val)
	}

	a2 := a
	fmt.Println("=======================")
	a2[5] = 100
	fmt.Printf("a2[%d]=%d\n",5,a2[5]) // a2[5]=100
	fmt.Printf("a[%d]=%d\n",5,a[5]) // a[5]=857

	fmt.Println("=======================")
	var c [5]int
	modify(c)
	for key,val := range c {
		fmt.Printf("c[%d]=%d\n",key,val)
	}
	fmt.Println("=========modify1==============")
	modify1(&c)
	fmt.Printf("c[%d]=%d\n",0,c[0])

	// 数组初始化
	fmt.Println("=========数组初始化=========")
	var age0 [5]int = [5]int{1,2,3,4,5}
	fmt.Println(age0)

	var age1 = [5]int{17,18,19,20,21}
	fmt.Println(age1)

	var age2 = [...]int{17,18,19,20}
	fmt.Println(age2)

	var strArr = [5]string{4:"ll",3:"ooo"}
	fmt.Println(strArr)

	// 多维数组
	fmt.Println("=========多维数组=========")
	var age3 [5][3]int
	fmt.Println(age3)

	var age4 = [...][2]int{{7,8},{13,14},{5,6}}
	fmt.Println(age4)
	for _,v := range age4 {
		fmt.Println(v)
		for _,v2:=range v {
			fmt.Println(v2)
		}
	}
}
func modify(arr [5]int)  {
	arr[0] =100
	return
}

func modify1(arr *[5]int)  {
	(*arr)[0] = 88
	return
}