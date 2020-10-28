package main

import (
	"flag"
	"fmt"

	"github.com/go-ini/ini"
	"github.com/labstack/gommon/log"
	"github.com/linzhenlong/my-go-code/2020/logagent/agent2/conf"
)

var (
	confPath *string
	cfg      = &conf.AppConf{}
)

func main() {
	//初始化获取配置信息.
	flag.Parse()
	err := ini.MapTo(cfg, *confPath)
	if err != nil {
		log.Fatalf("配置文件加载失败err:%v", err)
	}
	fmt.Println(cfg)

}

func init() {
	confPath = flag.String("conf", "./config.ini", "配置文件")
}
