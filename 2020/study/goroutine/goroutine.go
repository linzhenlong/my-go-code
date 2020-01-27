package main
import(
	"time"
	_"runtime"
	"fmt"

)

func main() {
	var a [10]int
	for i:=0;i<10;i++ {
		go func(i int) {
			for {
				//fmt.Printf("hello from goroutine %d\n", i)
				a[i] ++ 
				//runtime.Gosched() // 
			}	
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}