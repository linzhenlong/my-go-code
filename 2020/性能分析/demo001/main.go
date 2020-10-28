package main

import "fmt"

import "os"

import "runtime/pprof"

import "time"


// 写一个用于测试的方法

func myTest01() {
	var c chan int
	for {
		select {
		case v := <-c:
			fmt.Printf("chan read v:%d\n",v)
		default:
			// 先不写
			//time.Sleep(time.Second)
		}
	}
}

// 分析命令 go tool pprof main.go ./cpu.pprof
func main() {
	cpuPprof, _ := os.OpenFile("./cpu2.pprof",os.O_CREATE|os.O_RDWR,0666)
	defer cpuPprof.Close()

	// cpu 分析
	pprof.StartCPUProfile(cpuPprof)
	defer pprof.StopCPUProfile()
	for i:=0;i<8;i++ {
		go myTest01()
	}
	
	time.Sleep(time.Second * 15)
}