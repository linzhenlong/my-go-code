package etcd

import (
	"context"
	"encoding/json"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

var (
	etcdClient *clientv3.Client
)

// LogEntry ...
type LogEntry struct {
	Path  string `json:"path"`
	Topic string `json:"topic"`
}

// Init etcd 初始化.
func Init(address []string, timeout time.Duration) (err error) {
	etcdClient, err = clientv3.New(clientv3.Config{
		Endpoints:   address,
		DialTimeout: time.Millisecond * timeout,
	})
	return
}

// GetConf 从etcd中根据key获取配置项
func GetConf(key string) (logEntrys []*LogEntry, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	resp, err := etcdClient.Get(ctx, key)
	cancel()
	if err != nil {
		log.Printf("etced get err:%s", err.Error())
		return
	}
	for _, ev := range resp.Kvs {
		err = json.Unmarshal(ev.Value, &logEntrys)
		if err != nil {
			log.Printf("Unmarshal err:%s", err.Error())
			return
		}
	}
	return
}

// WatchLogConf ...
func WatchLogConf(key string, ch chan<- []*LogEntry) {
	watchChan := etcdClient.Watch(context.Background(), key)
	for wresp := range watchChan {
		for _, evt := range wresp.Events {
			log.Printf("Type:%v,key:%v,value:%v\n", evt.Type, evt.Kv.Key, evt.Kv.Value)
			newConf := make([]*LogEntry, 0)
			if evt.Type != clientv3.EventTypeDelete {
				// 判断一下非删除操作.
				err := json.Unmarshal(evt.Kv.Value, &newConf)
				if err != nil {
					log.Printf("WatchLogConf err :%v\n", err)
					continue
				}
			}
			log.Printf("WatchLogConf get newconf :%v\n", newConf)
			ch <- newConf
		}
	}
}
