package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CharCount struct {
	ChCount int // 英文的个数
	NumCount int // 数字的个数
	SpaceCount int // 空格的个数
	OtherCount int // 其他字符的个数
}

func main()  {
	// 统计一个文件中，英文，数字，空格及其他字符的个数

	var charCount CharCount

	fileName := "/Users/smzdm/Desktop/abc.txt"

	file ,err := os.Open(fileName)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		str ,err := reader.ReadString('\n')

		for _,v := range str {
			fmt.Println(string(v))
			switch  {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				charCount.ChCount++
			case v >= '0' && v<='9':
				charCount.NumCount++
			case v == ' ' || v == '\t':
				charCount.SpaceCount++
			default:
				charCount.OtherCount++
			}
		}
		if err == io.EOF {
			break
		}
	}
	fmt.Printf("英文:%d\n",charCount.ChCount)
	fmt.Printf("数字:%d\n",charCount.NumCount)

}

