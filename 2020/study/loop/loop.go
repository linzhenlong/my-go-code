package main

import "fmt"

import "strconv"

import "os"

import "bufio"

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
	scanner := bufio.NewScanner(file)
	// 省略递增条件
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	defer file.Close()
}

// for 省略结束条件
func forever() {
	for {
		fmt.Println("死循环")
	}
}
func main() {
	fmt.Println(
		convertTobin(5), // 101
		convertTobin(13),
	)

	filename := "abc.txt"
	printFile(filename)
	forever()
}