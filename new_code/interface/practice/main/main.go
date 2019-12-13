package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

// 1.声明Hero结构体
type Hero struct {
	Name string
	Age int
}
// 2.声明Hero结构体切片
type HeroSlice []Hero

// 3.实现Sort - Interface 接口
func (hs HeroSlice)Len() int {
	return len(hs)
}
// Less方法就是决定你使用哪种方式排序
// 1.按Hero的年龄排序从小到大排序
func (hs HeroSlice) Less(i,j int) bool  {
	return hs[i].Age < hs[j].Age
}

func (hs HeroSlice) Swap(i, j int)   {
	temp := hs[i]
	hs[i] = hs[j]
	hs[j] = temp
}

type Student struct {
	Name string
	Age int
	Score float64
}

type StuSlice []Student

func (s StuSlice)Len() int {
	return len(s)
}
func (s StuSlice)Less(i, j int) bool {
	return s[i].Score > s[j].Score
}
func (hs StuSlice) Swap(i, j int) {
	// 两个变量交互的简便写法
	hs[i], hs[j] = hs[j], hs[i]
}


func main()  {
	// 先定义一个数组切片
	var intSlice = []int{0,-1,10,9,8,66,109,22,4,6}

	sort.Ints(intSlice)
	fmt.Println(intSlice)

	// 测试我们是否可以对结构体切片进行排序

	var heroes HeroSlice
	for i:=0;i<10;i++ {
		hero := Hero{
			Name:fmt.Sprintf("英雄%d",i),
			Age:rand.Intn(80),
		}
		heroes = append(heroes,hero)
	}
	// 排序前
	fmt.Println("======排序前=======")
	for _,v := range heroes {
		fmt.Printf("名称:%s，年龄:%d\n",v.Name,v.Age)
	}
	fmt.Println("======排序后=======")
	sort.Sort(heroes)
	for _,v := range heroes {
		fmt.Printf("名称:%s，年龄:%d\n",v.Name,v.Age)
	}

	var students StuSlice
	for i:=0;i<10;i++ {
		score := math.Floor(rand.Float64() * 100)
		stu := Student{
			Name:fmt.Sprintf("学生%d",rand.Intn(100)),
			Age:rand.Intn(30),
			Score:score ,
		}
		students = append(students, stu)
	}
	fmt.Println("======学生排序前=======")
	for _,v := range students {
		fmt.Printf("名称:%s，分数:%.2f\n",v.Name,v.Score)
	}
	sort.Sort(students)
	fmt.Println("======学生排序后=======")
	for _,v := range students {
		fmt.Printf("名称:%s，分数:%.2f\n",v.Name,v.Score)
	}
}
