package main

import (
	"bufio"
	"fmt"
	"os"
)

// ValNode 数据节点.
type ValNode struct {
	row int
	col int
	val int
}

func main() {
	// 1.定义一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 // 黑棋
	chessMap[2][3] = 2 // 白棋

	// 2.输出原始数组
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		// 换个行.
		fmt.Println()
	}
	// 3.转成稀疏数组
	// (1).遍历chessMap,如果我们发现有一个元素的值不为0
	// 那么我们创建一个node结构体
	// (2).将其放到对应的切片里.

	var sparseArr []ValNode
	// 标准的稀疏数组应该还有一个,记录元素的二维数组的规模(行跟列，默认值)
	defaultNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr, defaultNode)
	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				node := ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, node)
			}
		}
	}
	fmt.Printf("%#v\n", sparseArr)

	file, err := os.OpenFile("./chessmap.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	defer file.Close()
	w := bufio.NewWriter(file)

	// 输出稀疏数组
	for i, nodeVal := range sparseArr {
		fmt.Printf("%d: %d\t %d\t %d\t \n", i, nodeVal.row, nodeVal.col, nodeVal.val)
		outStr := fmt.Sprintf("%d\t %d\t %d\t \n", nodeVal.row, nodeVal.col, nodeVal.val)
		w.Write([]byte(outStr))
	}

	w.Flush()

}
