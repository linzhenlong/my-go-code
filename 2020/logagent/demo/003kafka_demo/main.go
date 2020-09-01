package main

import "github.com/Shopify/sarama"

import "fmt"

func main() {

	config := sarama.NewConfig()
	//config.Version = sarama.V0_8_2_0

	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092", "127.0.0.1:9093"}, config)
	defer consumer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	partitionList, _ := consumer.Partitions("nginx-log")
	for _, p := range partitionList {
		msg, err := consumer.ConsumePartition("nginx-log", p, sarama.OffsetNewest)
		if err != nil {
			fmt.Println(err)
			continue
		}
		for {
			data := <-msg.Messages()
			fmt.Printf("Partition:%d,Offset:%d,key:%s,value:%s\n", data.Partition, data.Offset, string(data.Key), string(data.Value))
		}
	}
}
