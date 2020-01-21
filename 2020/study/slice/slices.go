package main

import "fmt"

// 切片是引用类型.
func updateSlice(s []int) {
	s[0] = 100
}
func printSlice(s []int) {
	fmt.Printf("s=%v,len=%d,cap=%d\n", s, len(s), cap(s))
}

func main() {
	arr := [...]int{0,1,2,3,4,5,6,7}
	s := arr[2:6]
	fmt.Println("arr[2:6]=",s)
	s1 := arr[:6]
	fmt.Println("arr[:6]=",s1)
	s2 := arr[2:]
	fmt.Println("arr[2:]=",s2)
	fmt.Println("arr[:]=",arr[:])

	updateSlice(s1)
	fmt.Println("arr[:6] ,s1 =",s1)
	updateSlice(s2)
	fmt.Println("arr[:6] ,s2 =",s2)
	
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	fmt.Println("切片的扩展")

	arr2 := [...]int{0,1,2,3,4,5,6,7,8,9}
	fmt.Println("arr2 ", arr2)
	s3 := arr2[2:6]
	fmt.Printf("arr2[2:6](s3)=%v,len(s3)=%d,cap(s3)=%d]\n",s3,len(s3),cap(s3))
	s4 := s3[3:5] // 5,6
	fmt.Printf("s3[3:5](s4)=%v,len(s4)=%d,cap(s4)=%d\n",s4,len(s4),cap(s4))
	
	// slice添加元素时，如果超月cap时，系统会重新分配更大的的底层数组
	// 由于值传递的关系，必须接手append的值 
	s5 := append(s4, 100) 
	fmt.Printf("s5 := append(s4, 100)=%v len(s5)=%d,cap(s5)=%d \n", s5, len(s5), cap(s5))

	// 创建一个slice
	fmt.Println("creating slice....")
	// 创建:方式1
	//var s []int // 默认值是nil
	// 创建:方式2
	//s := []int{2,4,6,8} // 默认值[2,4,6,8]
	// 创建:方式3
	sNew := make([]int,3,4) // 默认值[0,0,0]
	fmt.Println("默认值", s)
	if s == nil {
		fmt.Println("s的默认值是nil")
	} else {
		fmt.Println("s的默认值不是nil")
	}
	fmt.Println(sNew)

	// 添加元素
	for i:=0;i<100;i++ {
		sNew = append(sNew, i*2 + 1)
		printSlice(sNew)
	}
	fmt.Println("snew==>",sNew)
	// 拷贝
	fmt.Println("copying slice.....")
	 sNew1 := []int{2,4,6,8}
	copy(sNew,sNew1)
	printSlice(sNew1)
	printSlice(sNew)

	// 删除
	fmt.Println("deleting elements from slice")
	// 将sNew 中的8给删掉
	sNew = append(sNew[:3],sNew[4:]...)
	printSlice(sNew)
	
	fmt.Println("删除切片前面的元素")
	//font := sNew[0]
	sNew = sNew[1:]
	printSlice(sNew)
	
	fmt.Println("删除切片尾部的元素")
	//tail := sNew[len(sNew)-1]
	sNew = sNew[:len(sNew)-1]
	printSlice(sNew)
}