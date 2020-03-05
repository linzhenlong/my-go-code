package main

import(
	"runtime"
	"fmt"

)

func main() {
	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num)
	fmt.Println(num)
}