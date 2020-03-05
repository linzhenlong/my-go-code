package main

import "fmt"

// 获取最长不重复子串长度.
/* func lengthOfNoRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		//fmt.Println(lastOccurred)
		lastI, ok := lastOccurred[ch]
		if ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
} */

var  lastOccurred = make([]int, 0xffff)

func lengthOfNoRepeatingSubStr(s string) int {
	//lastOccurred := make([]int, 0xffff)
	// 运行之前将 将lastOccurred 数据清空.
	for i := range lastOccurred {
		lastOccurred[i] = 0
	}
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		//fmt.Println(lastOccurred)
		lastI := lastOccurred[ch]
		if lastI > start {
			start = lastI
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i+1
	}
	return maxLength
}

func main() {

	fmt.Println(lengthOfNoRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNoRepeatingSubStr("bbbbb"))
	fmt.Println(lengthOfNoRepeatingSubStr("pwwkew"))
	fmt.Println(lengthOfNoRepeatingSubStr(""))
	fmt.Println(lengthOfNoRepeatingSubStr("b"))
	fmt.Println(lengthOfNoRepeatingSubStr("abcdefg"))
	fmt.Println(lengthOfNoRepeatingSubStr("黑化肥发黑会回发"))

}
