package main

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/go-ini/ini"
	"github.com/linzhenlong/my-go-code/2020/logagent/agent/conf"
	"github.com/linzhenlong/my-go-code/2020/logagent/agent/etcd"
	"github.com/linzhenlong/my-go-code/2020/logagent/agent/kafka"
	"github.com/linzhenlong/my-go-code/2020/logagent/agent/taillog"
	"github.com/linzhenlong/my-go-code/2020/logagent/agent/utils"
)

var (
	cfg = &conf.AppConf{}
)

// func run() {
// 	// 1. 读日志
// 	for {
// 		select {
// 		case line := <-taillog.ReadLog():
// 			partition, offset, err := kafka.SendToKafka(cfg.Topic, line.Text)
// 			if err != nil {
// 				fmt.Println("SendToKafka err", err)
// 				break
// 			}
// 			fmt.Printf("SendToKafka success partition:%d,offset:%d\n", partition, offset)
// 		default:
// 			time.Sleep(time.Second)
// 		}
// 	}
// }

func main() {
	// 0.加载配置文件.
	/* cfg, _ := ini.Load("./conf/config.ini")
	kafkAddrs := cfg.Section("kafka").Key("address").Strings(",")
	//topic := cfg.Section("kafka").Key("topic").String()
	logFile := cfg.Section("taillog").Key("logfile").String() */
	ini.MapTo(cfg, "./conf/config.ini")

	// 1.初始化kafka连接
	err := kafka.InitKafka(cfg.KafkaConf.Address, cfg.KafkaConf.ChanMaxSize)
	if err != nil {
		fmt.Printf("InitKafka err:%v\n", err)
		return
	}
	// 2.初始化etcd
	err = etcd.Init(cfg.EtcdConf.Address, cfg.EtcdConf.TimeOut)
	if err != nil {
		fmt.Printf("ectd init err:%v\n", err)
		return
	}
	ip, err := utils.GetOutboundIP()
	if err != nil {
		ip = "127.0.0.1"
	}
	cfg.EtcdConf.CollectLogkey = strings.ReplaceAll(cfg.EtcdConf.CollectLogkey, "%s", ip)
	fmt.Println(cfg.EtcdConf.CollectLogkey)
	// 2.1 从etcd 中获取日志收集项的配置信息
	logEntryConf, err := etcd.GetConf(cfg.EtcdConf.CollectLogkey)
	if err != nil {
		log.Fatalf("etcd.GetConf err:%v", err)
	}
	fmt.Println(logEntryConf)
	for _, value := range logEntryConf {
		fmt.Println(value)
	}
	taillog.Init(logEntryConf)
	wg := sync.WaitGroup{}
	wg.Add(1)
	// 2.2 派一个哨兵，去监视日志收集项
	go etcd.WatchLogConf(cfg.EtcdConf.CollectLogkey, taillog.NewConfChan())
	wg.Wait()
	taillog.Init(logEntryConf)
	// err = taillog.InitTail(cfg.LogFile)
	// if err != nil {
	// 	fmt.Printf("InitTail err:%v", err)
	// 	return
	// }
	// run()

}
