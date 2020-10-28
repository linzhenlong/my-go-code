package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/axgle/mahonia"
)

// 读数据

func read() {
	dataBytes, err := ioutil.ReadFile("./data/kaifang.txt")
	if err != nil {
		panic(err)
	}
	dataStr := string(dataBytes)
	dataSlice := strings.Split(dataStr, "\n\r")
	for _, value := range dataSlice {
		newValue := convertEncoding(value, "GBK")
		fmt.Println(newValue)
	}

}

// 缓冲读取...
func read2() {
	file, err := os.Open("../data/kaifang.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		lineBytes, _, err := reader.ReadLine()
		if err == io.EOF {
			 break
		}
		if err != nil {
			continue
		}
		line := convertEncoding(string(lineBytes), "GBK")
		fmt.Println(line)
	}
}

// gbk 转utf-8
func convertEncoding(srcStr, encoding string) (dstStr string) {
	decoder := mahonia.NewDecoder(encoding)
	// 转utf8
	dstStr = decoder.ConvertString(srcStr)
	return
}

func main() {
	read2()
}
