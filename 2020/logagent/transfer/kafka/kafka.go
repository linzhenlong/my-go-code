package kafka

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Shopify/sarama"
	"github.com/linzhenlong/my-go-code/2020/logagent/transfer/conf"
	"github.com/linzhenlong/my-go-code/2020/logagent/transfer/elasticsearch"
)

// Init kafka
func Init(conf conf.KafkaConf) error {
	config := sarama.NewConfig()
	config.Net.DialTimeout = time.Duration(conf.DialTimeout) * time.Microsecond

	if len(conf.Address) == 0 {
		return errors.New("kafka address 不能为空")
	}
	if conf.Topic == "" {
		return errors.New("kafka Topic 不能为空")
	}
	consumer, err := sarama.NewConsumer(conf.Address, config)
	log.Printf("partitionList :%#v,%v", conf, err)
	if err != nil {
		return err
	}
	// 获取分区列表
	partitionList, err := consumer.Partitions(conf.Topic)
	log.Printf("partitionList :%v", partitionList)
	if err != nil {
		return err
	}
	for _, partiton := range partitionList {
		partitionConsumer, err := consumer.ConsumePartition(conf.Topic, partiton, sarama.OffsetNewest)
		if err != nil {
			log.Printf("ConsumePartition err :%v\n", err)
			continue
		}
		log.Printf("ConsumePartition err :%v\n", partitionConsumer)
		//defer partitionConsumer.AsyncClose()
		go readKakfa(partitionConsumer)
	}
	return nil
}

func readKakfa(partitionConsumer sarama.PartitionConsumer) {
	for msg := range partitionConsumer.Messages() {
		fmt.Printf("Partition:%d,topic:%s,offset:%d,key:%v,value:%v\n", msg.Partition, msg.Topic, msg.Offset, msg.Key, string(msg.Value))
		log := &elasticsearch.LogData{
			Msg:       string(msg.Value),
			Topic:     msg.Topic,
			Offset:    msg.Offset,
			Partition: msg.Partition,
		}
		elasticsearch.SendToChan(log)
	}
}
