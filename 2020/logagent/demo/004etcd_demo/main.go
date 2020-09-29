package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {
	fmt.Println("HAHAHAH")
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:23791", "127.0.0.1:23792", "127.0.0.1:23793"},
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		log.Printf("connect etcd err:%v", err)
		return
	}

	defer client.Close()
	// put
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	for i := 0; i < 100; i++ {
		key := "lzl-" + strconv.Itoa(i)
		value := "hello" + strconv.Itoa(i)
		_, err = client.Put(ctx, key, value)
		if err != nil {
			log.Printf("put err:%v", err)
			continue
		}
	}
	/*jsonStr := `[
		{
			"path":"/Users/smzdm/webroot/my-docker/docker-file/docker-compose/nginx-tomcat/logs/access.log",
			"topic":"nginx-log"
		},
		{
			"path":"/Users/smzdm/webroot/my-docker/docker-file/docker-compose/nginx-tomcat/logs/access1.log",
			"topic":"web-log"
		},
		{
			"path":"/Users/smzdm/webroot/my-docker/docker-file/docker-compose/nginx-tomcat/logs/access2.log",
			"topic":"web-log2"
		}
	]`*/
	jsonStr := `[
		{
			"path":"/Users/smzdm/webroot/my-docker/docker-file/docker-compose/nginx-tomcat/logs/access.log",
			"topic":"nginx-log"
		}
	]`
	ip, _ := GetOutboundIP()
	etcdKey := fmt.Sprintf("/logagent/%s/collect_config", ip)
	client.Put(ctx, etcdKey, jsonStr)
	//client.Delete(ctx, "/logagent/collect_config")
	cancel()
	// get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp2, err := client.Get(ctx, "lzl", clientv3.WithPrefix())
	if err != nil {
		log.Printf("get err:%v", err)
	}
	for _, val := range resp2.Kvs {
		log.Printf("%s:%s\n", val.Key, val.Value)
	}
}

// GetOutboundIP ...
func GetOutboundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	//fmt.Println(localAddr.String())
	ip = strings.Split(localAddr.IP.String(), ":")[0]
	return
}
