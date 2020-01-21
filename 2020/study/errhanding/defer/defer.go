package main
import(
	"fmt"
	"errors"
	"bufio"
	"os"
	"github.com/linzhenlong/my-go-code/2020/study/errhanding/defer/fib"

)

// 确保调用在函数结束时发生
// 参数在defer语句时计数
// defer 列表为先进后出.
func tryDefer() {
	for i:=0;i<100;i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}
	}
}

func writeFile(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)

	// err = errors.New("this is custom error") // 自定义error
	if err != nil {
		if pathError , ok := err.(*os.PathError);!ok {
			panic(err)
		} else {
			fmt.Printf("%s,%s,%s \n",
				pathError.Op,
				pathError.Path,
				pathError.Error(),
			)
		}
		return
		//fmt.Println("error:", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := fib.Fibonacci()
	for i:=0;i<20;i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	//tryDefer()
	
	writeFile("fib.txt")
}