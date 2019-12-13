package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main()  {

	file :="/Users/smzdm/Desktop/10.255.19.90-2.txt"

	// 打开并创建一个新文件
	// f , err := os.OpenFile(file, os.O_RDWR | os.O_CREATE , 0777)

	// 打开一个存在的文件并清空内容
	f , err := os.OpenFile(file, os.O_RDWR | os.O_APPEND , 0777)
	if err !=nil {
		fmt.Println(err)
		return
	}
	defer f.Close()


	reader := bufio.NewReader(f)
	for {
		content, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Println(content)
	}
	writer := bufio.NewWriter(f);
	for i:=0;i<10;i++ {
		writeStatus,err := writer.WriteString("fuck you \n")
		if err !=nil {
			fmt.Printf("%s", err)
			break
		}
		fmt.Println(writeStatus)
	}
	//因为write是代缓存的，所以写入之后还没有真正写到文件中去,因此需要flush 方法将内容写到磁盘
	writer.Flush()



	//该看视频248喽
}
