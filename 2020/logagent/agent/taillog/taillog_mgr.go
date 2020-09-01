package taillog

import (
	"github.com/linzhenlong/my-go-code/2020/logagent/agent/etcd"
	"log"
	"time"
)

var taskMgr *tailMgr

type tailMgr struct {
	logEntryList []*etcd.LogEntry
	taskMap      map[string]*TailTask
	newConfChan  chan []*etcd.LogEntry
}

// Init ...
func Init(logEntryConf []*etcd.LogEntry) {
	taskMgr = &tailMgr{
		logEntryList: logEntryConf,
		taskMap:      make(map[string]*TailTask, 32),
		newConfChan:  make(chan []*etcd.LogEntry), // 无缓冲区channel
	}
	for _, logEntry := range logEntryConf {
		task := NewTailTask(logEntry.Path, logEntry.Topic)
		mk := logEntry.Path + "_" + logEntry.Topic
		taskMgr.taskMap[mk] = task
	}
	go taskMgr.run()
}

// 监听自己的newConfChan ，有了新的配置过来就做对应的处理
func (t *tailMgr) run() {
	for {
		select {
		case newConf := <-t.newConfChan:
			// 1.新增配置
			// 2.删除配置
			// 3.配置变更
			for _, conf := range newConf {
				_, ok := t.taskMap[conf.Path+"_"+conf.Topic]
				// 说明这个配置存在不需要操作.
				if ok {
					continue
				}
				// 新增
				task := NewTailTask(conf.Path, conf.Topic)
				t.taskMap[conf.Path+"_"+conf.Topic] = task
			}
			//找出原来t.taskMap中有，但是newConf 中没有的

			for _, c1 := range t.logEntryList {
				isDel := true
				for _, c2 := range newConf {
					if c1.Path == c2.Path && c1.Topic == c2.Topic {
						isDel = false
						continue
					}
					if isDel {
						t.taskMap[c1.Path+"_"+c1.Topic].cancelFunc()
					}
				}
			}
			log.Println("新的配置来了", newConf)
		default:
			time.Sleep(time.Second)
		}
	}
}

// NewConfChan  一个函数，向外暴露taskMgr的newConfChan
func NewConfChan() chan<- []*etcd.LogEntry {
	return taskMgr.newConfChan
}
