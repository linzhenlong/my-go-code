package main

import (
	"bufio"
	"runtime/pprof"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/axgle/mahonia"
)

const (
	fileName = "./data/kaifang.txt"
	goodFile = "./data/kaifang_good.txt"
	badFile  = "./data/kaifang_bad.txt"
)

// gbk 转utf-8
func convertEncoding(srcStr, encoding string) (dstStr string) {
	decoder := mahonia.NewDecoder(encoding)
	// 转utf8
	dstStr = decoder.ConvertString(srcStr)
	return
}

// 按照34个省份划分数据
// 1.创建34个省份，34个数据管道
// 2.读优质数据，写入对应省份值
// 3.把省份管道写入到对应的文件中 (开34个协程)

// Province 省份对象
type Province struct {
	ID       string      // 身份证的前两位
	Name     string      // 省分名称
	File     *os.File    // 省份对应的文件，如:北京.txt
	ChanData chan string // 对应的管道
}

var wg sync.WaitGroup

func main() {
	pprofFile,_ := os.OpenFile("./data/pprof",os.O_RDWR|os.O_CREATE,0666)
	
	defer pprofFile.Close()
	defer pprof.WriteHeapProfile(pprofFile)
	// 声明一个map 存放所有省事
	provinces := make(map[string]*Province)
	ps := []string{"北京市11", "天津市12", "河北省13",
		"山西省14", "内蒙古自治区15", "辽宁省21", "吉林省22",
		"黑龙江省23", "上海市31", "江苏省32", "浙江省33", "安徽省34",
		"福建省35", "江西省36", "山东省37", "河南省41", "湖北省42",
		"湖南省43", "广东省44", "广西壮族自治区45", "海南省46",
		"重庆市50", "四川省51", "贵州省52", "云南省53", "西藏自治区54",
		"陕西省61", "甘肃省62", "青海省63", "宁夏回族自治区64", "新疆维吾尔自治区65",
		"香港特别行政区81", "澳门特别行政区82", "台湾省83"}

	// 遍历省份,创建省份对象

	for _, value := range ps {
		province := &Province{}
		name := value[:len(value)-2] // 获取名称
		id := value[len(value)-2:]

		file, _ := os.OpenFile("./data/省份/"+name+".txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		defer file.Close()
		province.Name = name
		province.ID = id
		province.File = file
		province.ChanData = make(chan string, 1024)
		provinces[id] = province
	}

	// 遍历所有省份,写数据
	for _, p := range provinces {
		wg.Add(1)
		go writeFile(p)
	}

	// 读优质文件
	good, _ := os.Open(goodFile)
	defer good.Close()

	// 缓冲读取
	reader := bufio.NewReader(good)

	// 逐行读取
	for {
		lineBytes, _, err := reader.ReadLine()
		if err == io.EOF {

			// 关闭管道
			for _, p := range provinces {
				fmt.Println(p.Name,"close")
				close(p.ChanData)
			}
			break
		}
		if err != nil {
			continue
		}
		// 获取身份证
		lineStr := string(lineBytes)
		lineSlice := strings.Split(lineStr, ",")
		if len(lineSlice) < 2 {
			continue
		}
		cardID := lineSlice[1]
		// 获取身份证号前两位
		ID := cardID[0:2]
		if province, ok := provinces[ID]; ok {
			province.ChanData <- lineStr
		}
	}
	wg.Wait()
}

func writeFile(p *Province) {
	for lineStr := range p.ChanData {
		_, err := p.File.WriteString(lineStr + "\n")
		if err != nil {
			continue
		}
		//fmt.Println(p.Name,lineStr,"===>success")
	}
	wg.Done()
}
