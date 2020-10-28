package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

func main() {
	fileName := "/Users/smzdm/webroot/my-docker/docker-file/docker-compose/nginx-tomcat/logs/access.log"
	config := tail.Config{
		ReOpen: true, //重新打开,当日志切分时,重新打开
		Follow: true, //是否跟随.
		Location: &tail.SeekInfo{ // 从文件的哪个位置开始读
			Offset: 0,
			Whence: 2,
		},
		MustExist: false,
		Poll:      true,
	}
	tails, err := tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail.TailFile err:", err)
		return
	}
	var (
		msg *tail.Line
		ok  bool
	)
	for {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen filename:%s\n", tails.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Printf("msg:%s\n", msg.Text)
	}
}
