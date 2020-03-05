package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main()  {

	// 99乘法表
	cheng()

	
	// 完数
	process(1000)

	// 回文
	var str string  = "你好阿11阿好你1"
	ish,hstr := ishuiwen(str)
	if ish {
		fmt.Println(hstr+"is hui")
	} else {
		fmt.Println(hstr+"is not hui")
	}

	// 统计中英文，数字，空格个数

	reader := bufio.NewReader(os.Stdin)
	res,_,err := reader.ReadLine()
	if err != nil {
		fmt.Println("error",err)
		return
	}
	wordCount,spaceCount,numberCount,otherCount := tongji(string(res))
	fmt.Printf("wordcount:%d,spacecount:%d,numberCount:%d,otherCount:%d\n",wordCount,spaceCount,numberCount,otherCount)

	//
	res2,_,err2 := reader.ReadLine()
	if err2 != nil {
		fmt.Println("error2",err)
		return
	}
	strSlice := strings.Split(string(res2),"+")
	if len(strSlice) !=2 {
		fmt.Println("please input a+b")
	}
	fmt.Printf("%q\n",strSlice)

	strNum1 := strings.TrimSpace(strSlice[0])
	strNum2 := strings.TrimSpace(strSlice[1])
	fmt.Println(addStr(strNum1, strNum2))

}
func addStr(strNum1,strNum2 string) string  {

	if len(strNum1) == 0 && len(strNum2) == 0 {
		return "0"
	}
	var(
		index1 int = len(strNum1) - 1
		index2 int = len(strNum2) - 1
		left int
		res string
	)
	for index1 >= 0 && index2 >=0 {
		c1 := strNum1[index1] - '0'
		c2 := strNum2[index2] - '0'
		sum := int(c1) + int(c2) + left
		if sum >=10 {
			left = 1
		} else {
			left = 0
		}
		c3 := (sum % 10) + '0'
		res = fmt.Sprintf("%c%s",c3,res)
		//fmt.Println(res)
		index1--
		index2--
	}
	for index1 >=0 {
		c1 := strNum1[index1] - '0'
		sum := int(c1)  + left
		if sum >=10 {
			left = 1
		} else {
			left = 0
		}
		c3 := (sum % 10) + '0'
		res = fmt.Sprintf("%c%s",c3,res)
		index1--
	}

	for index2 >=0 {
		c1 := strNum1[index2] - '0'
		sum := int(c1)  + left
		if sum >=10 {
			left = 1
		} else {
			left = 0
		}
		c3 := (sum % 10) + '0'
		res = fmt.Sprintf("%c%s",c3,res)
		index2--
	}
	if left ==1 {
		res = fmt.Sprintf("1%s",res)
	}
	return res
}

func cheng()  {
	// 99乘法表
	for i:=1;i<=9;i++ {
		for j:=1;j<=i;j++ {
			fmt.Printf("%d*%d=%d ",j,i,i*j)
		}
		fmt.Println()
	}
}

func iswannumber(n int) bool {
	var sum int
	for i:=1;i<n ;i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return n==sum
}

func process(n int)  {
	for i:=1;i<=n;i++ {
		if iswannumber(i) {
			fmt.Printf("%d is wan number \n",i)
		}
	}
}

func ishuiwen(str string) (bool,string) {
	runes := []rune(str)
	var newstr string
	for i:=len(runes); i>0;i-- {
		newstr += string(runes[i-1])
	}
	return str == newstr,newstr
}

func tongji(str string) (wordCount ,spaceCount,numberCount ,otherCount int) {
	runes := []rune(str)
	for _,val := range runes{
		switch  {
		case val >= 'a' && val <='z',val >= 'A' && val <= 'Z':
			wordCount++
		case val == ' ':
			spaceCount++
		case val >= '0' && val <='9':
			numberCount++
		default:
			otherCount++
		}
	}
	return
}