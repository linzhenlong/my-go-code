package main

import "fmt"

import "strconv"

import "os"

import "bufio"

import "io"

import "strings"

// int 转二进制
func convertTobin(n int) string {
	result := ""
	// for省略起始条件
	for ;n>0;n/=2{
		//最低位
		lsb := n %2
		result = strconv.Itoa(lsb) + result
	}
	return result
}
func printFile(filename string) {
	file ,err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	printFileContents(file)
	defer file.Close()
}

// for 省略结束条件
func forever() {
	for {
		fmt.Println("死循环")
	}
}
func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	// 省略递增条件
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
func main() {
	fmt.Println(
		convertTobin(5), // 101
		convertTobin(13),
	)

	filename := "abc.txt"
	printFile(filename)
	s := `abcadadasdad"d"
	lllll
	dadasdasdasd`
	printFileContents(strings.NewReader(s))
	//forever()
}