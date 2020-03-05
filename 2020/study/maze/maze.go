package main

import(
	"os"
	"fmt"

)
const fileName = "maze.in"
func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err !=nil {
		panic(err)
	}
	defer file.Close()
	var row,col int
	fmt.Fscanf(file, "%d %d", &row, &col)
	maze := make([][]int, row) // 6行
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}
type point struct {
	i, j int
}

// 逆时针上左下右遍历.
var dirs = [4]point{
	{-1,0},{0,-1},{1,0},{0,1},
}
func (p point) add(r point) point{
	return point{p.i+r.i, p.j+r.j}
}
func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >=len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >=len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}
func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	queue := []point{start}
	for len(queue) >0 {
		cur := queue[0]
		queue = queue[1:] // 移除头

		if cur == end {
			break
		}

		for _, dir := range dirs {
			next := cur.add(dir)

			// maze at next is 0
			// and steps at next is 0
			// and next != start
			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}
			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}
			if next == start {
				continue
			}
			curSteps , _ := cur.at(steps)
			steps[next.i][next.j] = curSteps + 1
			queue = append(queue, next)
		}
	}
	return steps
}
func main() {
	maze := readMaze(fileName)
	
	steps := walk(maze, point{0,0},point{len(maze)-1, len(maze[0])-1})

	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d ", val)
		}
		fmt.Println()
	}
}