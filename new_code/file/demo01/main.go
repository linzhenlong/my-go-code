package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main()  {



	//打开一个文件
	file, err := os.Open("/Users/smzdm/Desktop/10.255.19.90.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileInfo , _ := file.Stat();
	name := fileInfo.Name();
	fmt.Println(name)


	// 函数退出时关闭文件，否则会有内存泄露
	defer file.Close()

	// 创建一个*Reader 并且是代缓冲的
	Reader := bufio.NewReader(file)
	for {
		str , err := Reader.ReadString('\n')
		if err == io.EOF{  //表示读到文件末尾
			break
		}
		if str == "\n" {
			continue
		}
		fmt.Print(str)
	}

}