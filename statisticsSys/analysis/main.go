package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

type cmdParams struct {
	logFilePath   string
	routineNum    int
	targetLogPath string
}

type digData struct {
	time  string
	url   string
	refer string
	ua    string
}

type urlData struct {
	data digData
	uid  string
}

type urlNode struct {
	//
}

type storageBlock struct {
	counterType  string
	storageModel string
	unode        urlNode
}

var log = logrus.New()

func init() {
	log.Out = os.Stdout
	log.SetLevel(logrus.DebugLevel)
}
func main() {

	// 获取参数
	logFilePath := flag.String("logFilePath", "/data/log/tongji_access.log", "日志文件")
	routineNum := flag.Int("routineNum", 5, "协程数")

	targetLogPath := flag.String("targetLogPath", "/tmp/go.log", "目标日志路径")
	flag.Parse()

	params := cmdParams{
		logFilePath:   *logFilePath,
		routineNum:    *routineNum,
		targetLogPath: *targetLogPath,
	}

	// 打日志

	logFd, err := os.OpenFile(params.targetLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0655)
	fmt.Println(err)
	if err == nil {
		log.Out = logFd
		defer logFd.Close()
	}

	log.Infoln("程序启动了....")
	log.Infof("params logFilePath=%s,routineNum=%d,targetLogPath=%s ", params.logFilePath, params.routineNum, params.targetLogPath)

	// 初始化channel ,用于数据传递

	var logChannel = make(chan string, params.routineNum*3)
	var pvChannel = make(chan urlData, params.routineNum)
	var uvChannel = make(chan urlData, params.routineNum)

	var storageChannel = make(chan storageBlock, params.routineNum)

	// 日志消费者
	go readFileLineByLine(params, logChannel)

	// 创建一组日志处理
	for i := 0; i < params.routineNum; i++ {
		go logConsumer(logChannel, pvChannel, uvChannel)
	}

	// 创建pv,uv 统计器

	go pvCounter(pvChannel, storageChannel)
	go uvCounter(uvChannel, storageChannel)
	// 其他类型的统计也类似 xxxCounter()

	// 创建存储器
	go dataStorage(storageChannel)

	// 暂时想让程序一直跑着.
	time.Sleep(time.Second * 1000)


}

func dataStorage(storageChannel chan storageBlock) {

}

func pvCounter(pvChannel chan urlData, storageChannel chan storageBlock) {

}
func uvCounter(uvChannel chan urlData, storageChannel chan storageBlock) {

}

func logConsumer(logChannel chan string, pvChannel, uvChannel chan urlData) {

}

func readFileLineByLine(params cmdParams, logChannel chan string) error{
	fd, err := os.Open(params.logFilePath)
	if err != nil {
		log.Warningf("readFileLineByLine os.Open(%s) error =%s",params.logFilePath,err.Error())
		return err
	}
	defer fd.Close()
	bufRead := bufio.NewReader(fd)
	count := 0
	for {
		line, err := bufRead.ReadString('\n')
		log.Infof("line:",line)
		if err != nil {
			if err == io.EOF {
				time.Sleep(time.Second * 3)
				log.Infof("readFileLineByLine wait readLine :%d", count)
			} else {
				log.Warningf("readFileLineByLine bufRead.ReadString error = %s",err.Error())
			}
		}
		logChannel <- line
		count++

		if count % (params.routineNum*1000) == 0 {
			log.Infof("readFileLineByLine line:%d", count)
		}
	}

	return nil
}
