package main

import (
	"log"
	"strconv"
	"time"

	rabbitmq "github.com/linzhenlong/my-go-code/ms/rabbitmq-test/RabbitMQ"
)

func main() {

	rabbitmq := rabbitmq.NewRabbitMQSimple("linzl-test")
	ticker := time.NewTicker(time.Millisecond)

	after := time.After(time.Second * 60)
	log.Printf("START......")
	for {
		select {
		case t := <-ticker.C:
			rabbitmq.PublishSimple("fuck you" + t.Format("2006-01-02 15:04:05") + "," + strconv.Itoa(t.Nanosecond()))
		case <-after:
			log.Printf("Over......")
			ticker.Stop()
			return
		}
	}
}
