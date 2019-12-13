package main

import (
	"fmt"
	"io/ioutil"
)

const WRITE_FILE_NAME =  "/Users/smzdm/Desktop/10.255.19.90.2.txt"
const READ_FILE_NAME  = "/Users/smzdm/Desktop/10.255.19.90.txt"

func main()  {

	content , err := ioutil.ReadFile(READ_FILE_NAME)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile(WRITE_FILE_NAME,content,0777)
	if err == nil {
		fmt.Println("success")
	}
}