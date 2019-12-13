package main

import (
	"fmt"
	"os"
)

func main()  {
	var args []string
	args = os.Args
	for i, v := range args {
		fmt.Printf("args[%d]=%s\n",i, v)
	}
}
