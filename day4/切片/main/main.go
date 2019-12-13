package main

import "fmt"

func main()  {


	var sliceA = []int{1,2,3,4,5}
	var sliceB = []int{1,2,3,4,5}

	sliceA = append(sliceA,100)
	sliceC := append(sliceA,sliceB...)
	fmt.Println(sliceC)

	// for range 变量切片
	for key,v := range sliceC {
		fmt.Printf("sliceC[%d]=%d\n",key,v)
	}

	// 切片的拷贝
	s1 := []int{1,2,3,4,5}
	s2 := make([]int,3)
	copy(s2,s1)
	fmt.Println("s2=",s2) // s2= [1 2 3]

	s3 :=[]int{1,2,3}
	s3 = append(s3,s2...)
	s3 = append(s3,4,5,6)
	fmt.Println("s3=",s3)

	// string 于slice 区别
	str := "世界你好啊！"
	str2 := str[0:5]
	str3 := str[6:]
	fmt.Println(str2)
	fmt.Println(str3)

	// 修改字符串内容
	strArr :=[]rune(str)
	strArr[0] = '0'
	fmt.Println(string(strArr))
}
