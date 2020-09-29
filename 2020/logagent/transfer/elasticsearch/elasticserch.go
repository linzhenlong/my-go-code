package elasticsearch

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/linzhenlong/my-go-code/2020/logagent/transfer/conf"
	"github.com/olivere/elastic"
	v5 "gopkg.in/olivere/elastic.v5"
)

var (
	// EsClient ...
	EsClient   *v5.Client
	Esv7Client *elastic.Client

	DataChan = make(chan *LogData, 1000)
)

type EsHttpTransport struct {
}
type LogData struct {
	Topic     string `json:"topic"`
	Partition int32  `json:"partition"`
	Msg       string `json:"msg"`
	Offset    int64  `json:"offset"`
}

// Init es...
func InitV5(conf conf.ElasticSearchConf) error {
	options := []v5.ClientOptionFunc{}
	options = append(options, v5.SetURL(conf.Address...))
	options = append(options, v5.SetSniff(false))
	httpClient := &http.Client{
		Timeout:   time.Nanosecond,
		Transport: new(EsHttpTransport),
	}
	options = append(options, v5.SetHttpClient(httpClient))

	EsClient, err := v5.NewClient(options...)
	if err != nil {
		return err
	}
	srv := EsClient.Search().Index("kibana_sample_data_ecommerce").Type("_doc")
	query := v5.NewMatchAllQuery()
	res, err := srv.Query(query).Do(context.Background())
	fmt.Println(res)
	return nil
}

func InitV7(conf conf.ElasticSearchConf) (*elastic.Client, error) {
	options := []elastic.ClientOptionFunc{}
	options = append(options, elastic.SetURL(conf.Address...))
	options = append(options, elastic.SetSniff(false))
	options = append(options, elastic.SetMaxRetries(0))
	httpClient := &http.Client{

		Transport: new(EsHttpTransport),
	}
	options = append(options, elastic.SetHttpClient(httpClient))
	var err error
	Esv7Client, err = elastic.NewClient(options...)
	if err != nil {
		return Esv7Client, err
	}
	go sendToEs()
	return Esv7Client, nil
}

// RoundTrip ... 重写一下
func (e *EsHttpTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	//log.Printf("RoundTrip : %v\n", request)

	httpClient := &http.Client{
		Timeout: time.Nanosecond,
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 100,
		},
	}
	res, err := httpClient.Transport.RoundTrip(request)
	return res, err
}

func SendToChan(log *LogData) {
	DataChan <- log
}

func sendToEs() {
	for {
		select {
		case logdata := <-DataChan:
			res, err := Esv7Client.Index().Index(logdata.Topic).Type("_doc").BodyJson(logdata).Do(context.TODO())
			if err != nil {
				log.Println(err)
			}
			log.Println(res)
		default:
			time.Sleep(1000 * time.Microsecond)
			//log.Printf("sleep")
		}
	}
}
