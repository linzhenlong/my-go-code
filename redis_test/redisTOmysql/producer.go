package main

import "flag"

import "fmt"

import "math/rand"

import "time"

import "encoding/json"


var ruleTypes = []string {
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
// genData 生成数据.
func genData(data chan Article,genChan chan int) {
	for i:=0;i<10000;i++ {
		rand.Seed(time.Now().UnixNano())
		ruleID := rand.Int31n(10000)

		rules := Rules{
			RuleID: ruleID,
			RuleType:ruleTypes[rand.Intn(len(ruleTypes))],
			RuleWord:Keywords[rand.Intn(len(Keywords))],
			AddTime:time.Now().Format("2006-01-02 15:04:05"),
		}
		// rulesBytes, err := json.Marshal(rules)
		/* if err != nil {
			continue
		} */
		article := Article{
			ArticleID:rand.Int31n(1000000),
			UserID:rand.Int31n(1000000),
			Rules:rules,
		}
		fmt.Println("genData",i)
		data <- article
	}
	genChan <- 1
}

func lpushData(data chan Article, finshed chan int) {	
	for {
		article, ok := <- data
		if !ok {
			break
		}
		res,err := json.Marshal(article)
		if err != nil {
			fmt.Println(err)
			break
		}
		n , err := redisClient.LPush("article_list",string(res)).Result()
		if err != nil {
			fmt.Println("lpushData LPush ERROR", err.Error())
			break
		}
		fmt.Println("lpushData",n)
	}
	finshed <- 1
}

func init() {
	redisClient = CreateRedisClient()
}

func main() {

	routineNum := flag.Int("routineNum", 5, "携程数")
	flag.Parse()
	fmt.Println("携程数",*routineNum)
	
	
	// 生成数据
	var dataChan = make(chan Article)
	var genChan = make(chan int, *routineNum)
	for i:=0;i<*routineNum;i++ {
		go genData(dataChan, genChan)
	}
	for i:=0;i<*routineNum;i++ {
		<- genChan
		if i == *routineNum-1 {
			close(dataChan)
		}
	}
	
	var finshed = make(chan int, *routineNum)
	// lpush DATA
	for i:=0;i<*routineNum;i++ {
		go lpushData(dataChan,finshed)
	}
	for i:=0;i<*routineNum;i++ {
			<- finshed
			fmt.Println(i)
	}
	defer redisClient.Close()
	fmt.Println("done")
	

}
