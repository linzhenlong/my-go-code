package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const WRITE_FILE_NAME  =  "/Users/smzdm/Desktop/10.255.19.90.2.txt"
const READ_FILE_NAME  = "/Users/smzdm/Desktop/10.255.19.90.txt"

func main()  {

	file , err := os.Create(WRITE_FILE_NAME)
	if err != nil {
		fmt.Println(err)
	}
	file2, err := os.Open(READ_FILE_NAME)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	defer file2.Close()

	reader := bufio.NewReader(file2)

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if str == "\n" {
			continue
		}
		if strings.Contains(str,"#") {
			fmt.Println(str)
			continue;
		}
		_, err = file.Write([]byte(str))
		if err !=nil {
			fmt.Println(err)
		}

	}
}