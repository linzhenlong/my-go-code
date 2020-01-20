package main

import "fmt"

// 获取最长不重复子串长度.
func lengthOfNoRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i ,ch := range []rune(s) {
		fmt.Println(lastOccurred)
		lastI, ok := lastOccurred[ch]
		if ok && lastI >= start {
			start = lastI + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
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