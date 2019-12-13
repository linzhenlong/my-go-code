package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
三个班，每个班有5个学生
求出每个班的平均分，及所有班级的平均分
 */
func main() {

	var students [3][5]int
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<len(students) ;i++  {
		for j:=0;j<len(students[i]);j++ {
			students[i][j] = rand.Intn(100)
		}
	}

	var (
		allClassSum float64
	)
	fmt.Println(students)
	for i:=0;i<len(students) ;i++  {
		var classSum float64;
		for j:=0;j<len(students[i]);j++ {
			classSum += float64(students[i][j])
		}
		allClassSum +=classSum
		fmt.Printf("班级%v的,总分%v,平均分是%v\n",i+1,classSum,classSum/float64(len(students[i])))
	}
	fmt.Printf("全部班级的,总分%v,平均分是%v\n",allClassSum,allClassSum/float64(len(students)))

}


