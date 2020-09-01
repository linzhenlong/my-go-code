package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"strconv"
)

// 基于sarama第三⽅库开发的kafka client
func main() {
	config := sarama.NewConfig()
	// tailf包使⽤
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Version = sarama.V0_8_2_0
	config.Producer.Return.Errors = true

	// 构造⼀个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "test-Topic"

	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092", "127.0.0.1:9093"}, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	defer client.Close()

	for i := 0; i < 100000000; i++ {
		msgStr := "this is a test log" + strconv.Itoa(i)
		msg.Value = sarama.StringEncoder(msgStr)
		// 发送消息
		pid, offset, err := client.SendMessage(msg)
		if err != nil {
			fmt.Printf("SendMessage err:%v\n", err)
			return
		}
		fmt.Printf("SendMessage success 分区:%d,偏移量:%d\n", pid, offset)
	}
}
