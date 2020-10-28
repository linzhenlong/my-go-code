package main

import "github.com/Shopify/sarama"

import "time"

import "fmt"

const (
	kafkaTopic = "nginx-log"
)

func main() {
	cfg := sarama.NewConfig()
	cfg.Net.DialTimeout = time.Second * 5
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092", "127.0.0.1:9093"}, cfg)
	if err != nil {
		fmt.Println(err)
		return
	}
	partitionList, err := consumer.Partitions(kafkaTopic)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(partitionList)
	for partition := range partitionList {
		pc, err := consumer.ConsumePartition(kafkaTopic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("pc", pc)
		defer pc.AsyncClose()
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
		}(pc)
	}
	time.Sleep(time.Second * 10)
}
