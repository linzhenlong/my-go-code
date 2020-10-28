package main

import (
	"context"
	"fmt"
	"github.com/go-ini/ini"
	"go.etcd.io/etcd/clientv3"
	"time"
)

var addrs []string
var timeout time.Duration

func init() {
	cnf, _ := ini.Load("./config.ini")
	addrs = cnf.Section("etcd").Key("addrs").Strings(",")
	timeout, _ = cnf.Section("etcd").Key("timeout").Duration()
}

func main() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   addrs,
		DialTimeout: time.Millisecond * timeout,
	})
	if err != nil {
		fmt.Printf("connect etcd err:%v\n", err)
		return
	}
	defer client.Close()
	// 安排一个哨兵，一直监视者，这个key的变化(新增，修改，)
	watchChan := client.Watch(context.Background(), "lzl", clientv3.WithPrefix())
	for watchResp := range watchChan {
		for _, watchVal := range watchResp.Events {
			fmt.Printf("Type:%v key:%v value:%v\n", watchVal.Type, string(watchVal.Kv.Key), string(watchVal.Kv.Value))
		}
	}
}
