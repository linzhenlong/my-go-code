package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/axgle/mahonia"
)

const (
	fileName = "../data/kaifang.txt"
	goodFile = "../data/kaifang_good.txt"
	badFile  = "../data/kaifang_bad.txt"
)

// gbk 转utf-8
func convertEncoding(srcStr, encoding string) (dstStr string) {
	decoder := mahonia.NewDecoder(encoding)
	// 转utf8
	dstStr = decoder.ConvertString(srcStr)
	return
}

func main() {

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// 创建优质文件
	good, err := os.OpenFile(goodFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer good.Close()
	// 创建劣质文件
	bad, err := os.OpenFile(badFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer bad.Close()

	// 读文件
	fileReader := bufio.NewReader(file)

	for {
		lineBytes, _, err := fileReader.ReadLine()

		// 说明读完了
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}

		// 读出来的内容的字符集是gbk的，需要转成utf-8的
		gbkLineStr := string(lineBytes)
		utf8Line := convertEncoding(gbkLineStr, "GBK")

		// 根据行数据取身份证.
		lineArr := strings.Split(utf8Line, ",")

		// 判断切片长度大于等于2,身份证号长度大于18
		if len(lineArr) >=2 && len(lineArr[1]) == 18 {
			// 写到goodfile 里
			_, err = good.WriteString(utf8Line+"\n")
			if err != nil {
				fmt.Println(err)
				continue
			}
		} else {
			_, err = bad.WriteString(utf8Line+"\n")
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
		fmt.Println(lineArr)
	}
}
