package main

import "fmt"

// https://cloud.tencent.com/developer/article/1479645

// 场景1 
func test1() {
	vals := []int{0,1,2,3}
	valptr := make([]*int,0)
	for _, v := range vals {
		valptr = append(valptr, &v)
	}
	for _,v := range valptr {
		fmt.Println("2==>",v) // 2==> 0xc0000140a0
		fmt.Println("3==>",*v) // 3==> 3
	}
	// 原因:
	// go runtime中for range循环只会为v分配一次内存，后续只是给v赋值；跟for的语义是一样一样的,
	// 看test2()就容易理解多了
}
func test2() {
	for i:=0;i<3;i++ {
		fmt.Println("&i=",&i,"i=",i)
	}
	/**
	&i= 0xc0000b4038 i= 0
	&i= 0xc0000b4038 i= 1
	&i= 0xc0000b4038 i= 2
	**/
}

// 场景2 - 在closure中捕获循环变量
func test3() {
	kvs := map[string]int{
		"zero":0,
		"one":1,
		"two":2,
	} 
	for _, v := range kvs {
		defer func(){
			fmt.Println("defer func test3",v)
		}()
	}

	// go vet
	// ./main.go:42:29: loop variable v captured by func literal
	// 分析 
	//closure(func literal)中捕获变量是通过引用的方式,因此也会像场景1中那样,
	// v的地址虽然没变,但是值会随着循环变化
	// 可以通过 test4() 或test5()的方式解决
}
func test4() {
	kvs := map[string]int{
		"zero":0,
		"one":1,
		"two":2,
	} 
	for _, v := range kvs {
		v := v 
		// 说明:v (v1) := v (v2) 括号是注释是显式的创建了一个v的副本,也叫v;  
		// 这里两个v的生命周期不同, v2的生命周期是整个for循环,
		// v1的生命周期是for循环中的一个循环,但是这里由于closure对于v1的引用,
		// 所以在一个循环结束后,v会发生逃逸,并不会被立即回收
		defer func(){
			fmt.Println("defer func test4",v)
		}()
	}
}
func test5() {
	kvs := map[string]int{
		"zero":0,
		"one":1,
		"two":2,
	} 
	for _, v := range kvs {
		// 把v作为closure的参数,通过golang的pass-by-value,隐式的创建了一个v的副本
		defer func(v int){
			fmt.Println("defer func test5",v)
		}(v)
	}
}

func main() {
	test1() 
	test2()
	test3()
	test4()
	test5()
}