package main

import (
	"fmt"
	"strings"
)

func main() {

	for i := 1; i <= 5; i++ {
		fmt.Println(strings.Repeat("A", i))
	}

	// for range
	str := "hello world ,中国"
	for i, v := range str {
		fmt.Printf("index[%d] val[%c] len[%d]\n", i, v, len([]byte(string(v))))
	}

	// break continue
	for key, val := range str {
		if key == 2 {
			continue
		}
		if key > 5 {
			break
		}
		fmt.Printf("%d,%c\n", key, val)
	}

	// goto 和label语句
	LABEL1:
		for i := 0; i <= 5; i++ {
			for j := i; j <= 5; j++ {
				if j == 4 {
					continue LABEL1
				}
				fmt.Printf("i is :%d,and j is %d \n", i, j)
			}
		}


	/*i := 0
	HERE:

		fmt.Println(i)
		i++
		if i == 5 {
			return
		}
		goto HERE*/

	j:=0
	for {
		if j>= 3 {
			break
		}
		fmt.Println("",j)
		j++
	}
}
