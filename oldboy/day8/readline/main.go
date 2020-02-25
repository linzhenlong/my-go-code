package main

import (
	"bufio"
	"io"
	"log"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./test.log")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	defer file.Close()
	
	for {
		contents, isPrefix, err := reader.ReadLine()

		if err == io.EOF {
			// 读到文件末尾, 退出
			break
		}
		if err != nil {
			log.Printf("readline err : %s", err.Error())
				continue
		}
		fmt.Println(string(contents),isPrefix)
	}
}
