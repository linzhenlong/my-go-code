package main

import (
	"log"

	"github.com/go-ini/ini"
	"github.com/linzhenlong/my-go-code/2020/logagent/transfer/conf"
	"github.com/linzhenlong/my-go-code/2020/logagent/transfer/elasticsearch"
	"github.com/linzhenlong/my-go-code/2020/logagent/transfer/kafka"
)

var (
	cfg = &conf.Conf{}
)

func main() {
	//0.加载配置文件
	err := ini.MapTo(cfg, "./conf/cfg.ini")
	if err != nil {
		log.Fatalf("load config.ini err:%v", err)
	}
	log.Printf("conf:%v", cfg)
	//1. 初始化
	err = kafka.Init(cfg.KafkaConf)
	if err != nil {
		log.Fatalf("kafka.Init err:%v", err)
	}
	//2.从kafka消费数据
	//3.写数据
	_, err = elasticsearch.InitV7(cfg.EsConf)

	if err != nil {
		log.Fatalf("elasticsearch.Init err:%v", err)
	}

	select {}
}
