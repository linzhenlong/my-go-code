package kafka

import (
	"github.com/Shopify/sarama"
	"log"
	"time"
)

type logData struct {
	topic string
	data  string
}

var (
	producerClient sarama.SyncProducer // 声明一个全局的生产者客户端.
	logDataChan    chan *logData
)

// InitKafka 初始化kafka.
func InitKafka(addrs []string, chanSize int) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	producerClient, err = sarama.NewSyncProducer(addrs, config)
	if err != nil {
		log.Printf("sarama.NewSyncProducer err:%s addrs:%v\n", err.Error(), addrs)
	}
	// 初始化logDataChan
	logDataChan = make(chan *logData, chanSize)
	// 开启goroutine 写kafka.
	go sendToKafka()

	return
}

// sendToKafka .
func sendToKafka() {

	for {
		select {
		case logdata := <-logDataChan:
			msg := &sarama.ProducerMessage{}
			msg.Topic = logdata.topic
			msg.Value = sarama.StringEncoder(logdata.data)

			// 发送到kafka
			partition, offset, err := producerClient.SendMessage(msg)
			log.Printf("topic:%s,partition:%d,offset:%d err:%v\n", msg.Topic, partition, offset, err)
		default:
			time.Sleep(time.Second)
		}
	}
}

// SendToChan ...
func SendToChan(topic, log string) {
	logData := &logData{
		topic: topic,
		data:  log,
	}
	logDataChan <- logData
}
