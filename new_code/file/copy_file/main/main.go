package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CopyFile(dstFileName string ,srcFileName string) (int64, error) {

	srcFile , err := os.Open(srcFileName)
	if err !=nil {
		return 0, err
	}
	reader := bufio.NewReader(srcFile)

	dstFile ,err := os.OpenFile(dstFileName,os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return  0, err
	}
	writer := bufio.NewWriter(dstFile)

	defer srcFile.Close()
	defer dstFile.Close()
	return io.Copy(writer, reader)
}

func main()  {

	srcFileName := "/Users/smzdm/Desktop/IMG_7279.JPG"
	dstFileName := "/Users/smzdm/Desktop/IMG_7279-2.JPG"

	n , err :=CopyFile(dstFileName, srcFileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)
}