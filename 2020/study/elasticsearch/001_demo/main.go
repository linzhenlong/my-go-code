package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/olivere/elastic"
	"github.com/olivere/elastic/config"
)

// Person ...
type Person struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}


func main() {
	sniff := false

	cfg := &config.Config{
		URL:   "http://localhost:9200",
		Sniff: &sniff,
	}
	client, err := elastic.NewClientFromConfig(cfg)
	if err != nil {
		log.Fatalf("es init err:%s", err.Error())
	}
	log.Println("connect es succss")

	for i := 1; i <= 1000; i++ {
		person := Person{
			ID:      i,
			Name:    fmt.Sprintf("name%d", i),
			Age:     i,
			Married: true,
		}
		res, err := client.Index().Id(strconv.Itoa(i)).Index("person").Type("stu").BodyJson(person).Do(context.TODO())
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println(res)
		fmt.Println()
	}

}
