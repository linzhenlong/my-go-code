package main

import "flag"

import "fmt"

import "math/rand"

import "time"

import "encoding/json"

import "sync"

var ruleTypes = []string{
	"brand",
	"mall",
	"tag",
	"keyword",
}

// Keywords 关键词列表
var Keywords = []string{
	"神加个",
	"京东",
	"天猫",
	"手机",
	"电脑",
	"亚马逊",
}
var wg sync.WaitGroup

// genData 生成数据.
func genData(data chan Article) {
	for j := 0; j < 5; j++ {
		for i := 0; i < 100; i++ {
			rand.Seed(time.Now().UnixNano())
			ruleID := rand.Int31n(10000)

			rules := Rules{
				RuleID:   ruleID,
				RuleType: ruleTypes[rand.Intn(len(ruleTypes))],
				RuleWord: Keywords[rand.Intn(len(Keywords))],
				AddTime:  time.Now().Format("2006-01-02 15:04:05"),
			}
			// rulesBytes, err := json.Marshal(rules)
			/* if err != nil {
				continue
			} */
			article := Article{
				ArticleID: rand.Int31n(1000000),
				UserID:    rand.Int31n(1000000),
				Rules:     rules,
			}
			fmt.Println("genData", i)
			data <- article
		}
	}
	close(data)
	//genChan <- 1
}

func lpushData(data chan Article) {

	for article := range data {
		res, err := json.Marshal(article)
			if err != nil {
				fmt.Println(err)
				break 
			}
			n, err := redisClient.LPush("lzl", string(res)).Result()
			if err != nil {
				fmt.Println("lpushData LPush ERROR", err.Error())
				break 
			}
			fmt.Println("lpushData", n)
	}
	defer wg.Done()
}

func init() {
	redisClient = CreateRedisClient()
}

func main() {

	routineNum := flag.Int("routineNum", 5, "携程数")
	flag.Parse()
	fmt.Println("携程数", *routineNum)

	// 生成数据
	var dataChan = make(chan Article)

	//var genChan = make(chan int, *routineNum)
	//for i := 0; i < *routineNum; i++ {
	go genData(dataChan)
	//}

	//close(dataChan)
	/* for i:=0;i<*routineNum;i++ {
		fmt.Println(i)
		<- genChan
		if i == *routineNum-1 {
			close(dataChan)
		}
	} */

	// lpush DATA
	wg.Add(*routineNum)
	for i := 0; i < *routineNum; i++ {
		go lpushData(dataChan)
	}
	wg.Wait()
	defer redisClient.Close()
	fmt.Println("done")

}
