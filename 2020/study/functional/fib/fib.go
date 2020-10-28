package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// 斐波那契数列
// 1,1,2,3,5,8,13,21.....
//     a, b
//        a, b
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int

func fibonacci2() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// 实现io.Read接口.
func (intgen intGen) Read(p []byte) (n int, err error) {
	next := intgen()
	if next > 1000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	// 省略递增条件
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
func main() {

	f := fibonacci()
	fmt.Println(f()) // 1
	fmt.Println(f()) // 1
	fmt.Println(f()) // 2
	fmt.Println(f()) // 3
	fmt.Println(f()) // 5
	fmt.Println(f()) // 8
	fmt.Println(f()) // 13

	/* f := fibonacci2()
	printFileContents(f) */

}
