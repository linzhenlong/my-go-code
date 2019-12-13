package main

import "fmt"

// 结构体
type Point struct {
	x int
	y int
}


type Rect struct {
	leftUp, rightDown Point
}

type Rect2 struct {
	leftUp, rightDown *Point // 指针
}
/**
	结构体的所有字段在内存中是连续的
 */
func main()  {

	r1 := Rect{Point{1,2},Point{3,4}}
	// r1 有四个整数int 型
	fmt.Printf("r1.lefUp.x的内存地址:%p \n",&r1.leftUp.x) // r1.lefUp.x的内存地址:0xc420016140
	fmt.Printf("r1.lefUp.y的内存地址:%p \n",&r1.leftUp.y) //r1.lefUp.y的内存地址:0xc420016148
	fmt.Printf("r1.rightDown.x的内存地址:%p \n",&r1.rightDown.x) // r1.rightDown.x的内存地址:0xc420016150
	fmt.Printf("r1.rightDown.y的内存地址:%p \n",&r1.rightDown.y) // r1.rightDown.y的内存地址:0xc420016158
	// 出输出结果可以看出，内存地址是连续的

	//让r2有两个*Point类型，这两个*Point类型的本身的地址也是连续的，但是他们指向的地址不一定是联系的
	r2 :=Rect2{&Point{10,20},&Point{30,40}}
	fmt.Printf("r2.lefUp本身的内存地址:%p \n",&r2.leftUp) //r2.lefUp本身的内存地址:0xc42000e1e0
	fmt.Printf("r2.rightDown本身的内存地址:%p \n",&r2.rightDown) //r2.rightDown本身的内存地址:0xc42000e1e8

	//他们指向的地址不一定是连续的，这个要看编译器在当时运行时是如何分配的
	fmt.Printf("r2.lefUp本身的内存地址:%p \n",r2.leftUp) //r2.lefUp本身的内存地址:0xc420084010
	fmt.Printf("r2.rightDown本身的内存地址:%p \n",r2.rightDown) //r2.rightDown本身的内存地址:0xc420084020

}
