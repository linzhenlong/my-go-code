package taillog

import (
	"context"
	"fmt"
	"time"

	"github.com/hpcloud/tail"
	"github.com/linzhenlong/my-go-code/2020/logagent/agent/kafka"
)

var (
	tailObj *tail.Tail
)

// TailTask 一个日志收集的任务
type TailTask struct {
	path       string
	topic      string
	instance   *tail.Tail
	ctx        context.Context
	cancelFunc context.CancelFunc
}

// NewTailTask ...
func NewTailTask(path, topic string) (task *TailTask) {
	ctx, cancel := context.WithCancel(context.Background())
	task = &TailTask{
		path:       path,
		topic:      topic,
		ctx:        ctx,
		cancelFunc: cancel,
	}
	task.init() // 根据日志路径打开响应日志
	return
}

// Init ...
func (t *TailTask) init() {
	config := tail.Config{
		ReOpen: true, //是否重新打开文件
		Follow: true, //是否跟随
		Location: &tail.SeekInfo{ // 从文件的那个位置开始读
			Offset: 0,
			Whence: 2,
		},
		MustExist: false, // 文件不存在不报错
		Poll:      true,
	}
	var err error
	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Println("tail.TailFile err", err)
		return
	}
	go t.run()
}

// ReadLog .
func (t *TailTask) ReadLog() <-chan *tail.Line {
	return t.instance.Lines
}

// Run ...
func (t *TailTask) run() {
	for {
		select {
		case <-t.ctx.Done():
			// 退出.
			fmt.Printf("tail task :%v结束退出了\n", t.path+"_"+t.topic)
			return
		case line := <-t.ReadLog():
			// 发送到kakfa
			//kafka.SendToKafka(t.topic, line.Text)
			//先把日志数据发送到一个channel 中
			// 在kafka 的包里，有一个单独的goroutine去取日志数据发到kafka
			kafka.SendToChan(t.topic, line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}
