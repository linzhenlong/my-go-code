package main

import (
	"fmt"
	"os"
)

const WRITE_FILE_NAME =  "/Users/smzdm/Desktops"
//const READ_FILE_NAME  = "/Users/smzdm/Desktop/10.255.19.90.txt"

func main()  {

	isExist , err := FileExists(WRITE_FILE_NAME)
	if err != nil {
		fmt.Println(err);
	}
	if isExist {
		fmt.Println("exists")
	} else {
		fmt.Println("no exists")
	}
}

func FileExists(file string) (bool, error)  {
	_, err := os.Stat(file)
	if err == nil {  // 如果os.Stat() 函数返回的的错误值为nil 则证明文件或是文件夹存在
		return true,nil
	}
	if os.IsNotExist(err) { // 如果os.Stat()函数返回的错误值使用os.IsNotExist()判断为true 证明文件或是文件加不存在
		return false, nil
	}
	return false,err // 如果os.Stat()函数返回的错误值 为其他类型，则不确定文件是否存在
}

