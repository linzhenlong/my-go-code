package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/mgutz/str"
	"github.com/sirupsen/logrus"
	"io"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const HANDLE_GIF  = " /dot.gif?"

type cmdParams struct {
	logFilePath   string
	routineNum    int
	targetLogPath string
}

type digData struct {
	time  string
	url   string
	refer string
	ua    string
}

type urlData struct {
	data digData
	uid  string
	unode urlNode
}

type urlNode struct {
	//
	unType string  // 首页，列表页，或是详情页
	unRequestId string // 资源id
	unUrl string  // 当前页面的url
	unTime string // 当前访问时间
}

type storageBlock struct {
	counterType  string
	storageModel string
	unode        urlNode
}

var log = logrus.New()

var redisClient *redis.Client

func init() {
	log.Out = os.Stdout
	log.SetLevel(logrus.DebugLevel)
	redisClient = createRedis()


}
func createRedis()*redis.Client  {
	redisCli := redis.NewClient(&redis.Options{
		Network:"tcp",
		Addr:"127.0.0.1:6379",
		DB:0,
		Password:"",
		PoolSize:10,

	})
	return redisCli
}

func main() {

	defer redisClient.Close()
	fmt.Println("tj start")
	// 获取参数
	logFilePath := flag.String("logFilePath", "/data/log/tongji_access.log", "日志文件")
	routineNum := flag.Int("routineNum", 5, "协程数")

	targetLogPath := flag.String("targetLogPath", "/tmp/go.log", "目标日志路径")
	flag.Parse()

	params := cmdParams{
		logFilePath:   *logFilePath,
		routineNum:    *routineNum,
		targetLogPath: *targetLogPath,
	}

	// 打日志

	logFd, err := os.OpenFile(params.targetLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0655)
	fmt.Println(err)
	if err == nil {
		log.Out = logFd
		defer logFd.Close()
	}

	log.Infoln("程序启动了....")
	log.Infof("params logFilePath=%s,routineNum=%d,targetLogPath=%s ", params.logFilePath, params.routineNum, params.targetLogPath)

	// 初始化channel ,用于数据传递

	var logChannel = make(chan string, params.routineNum*3)
	var pvChannel = make(chan urlData, params.routineNum)
	var uvChannel = make(chan urlData, params.routineNum)

	var storageChannel = make(chan storageBlock, params.routineNum)

	// 日志消费者
	go readFileLineByLine(params, logChannel)

	// 创建一组日志处理
	for i := 0; i < params.routineNum; i++ {
		go logConsumer(logChannel, pvChannel, uvChannel)
	}

	// 创建pv,uv 统计器

	go pvCounter(pvChannel, storageChannel)
	go uvCounter(uvChannel, storageChannel)
	// 其他类型的统计也类似 xxxCounter()

	// 创建存储器
	go dataStorage(storageChannel)

	// 暂时想让程序一直跑着.
	//time.Sleep(time.Second * 1000)


}

func dataStorage(storageChannel chan storageBlock) {
	for block := range storageChannel{
		prefix := block.counterType + "_"

		// 逐层添加，加洋葱皮的过程
		// 维度:天-小时-分钟
		// 层级:大分类-小分类-详情页
		setKeys := []string{
			prefix+"day_"+ getTime(block.unode.unTime, "day"),
			prefix+"hour_"+ getTime(block.unode.unTime, "hour"),
			prefix+"min_"+ getTime(block.unode.unTime, "min"),
			prefix+block.unode.unType+"_day_"+ getTime(block.unode.unTime, "day"),
			prefix+block.unode.unType+"_hour_"+ getTime(block.unode.unTime, "hour"),
			prefix+block.unode.unType+"_min_"+ getTime(block.unode.unTime, "min"),
		}

		rowId := block.unode.unRequestId

		for _, key := range setKeys {
			ret,err := redisClient.ZIncrBy(key, 1, rowId).Result()
			if err != nil || ret <= 0 {
				log.Fatalf("redisClient error= ", err)
			}
		}
	}
}

func pvCounter(pvChannel chan urlData, storageChannel chan storageBlock) {
	for data := range pvChannel {
		sItem := storageBlock{
			counterType:"pv",
			storageModel:"ZINCRBY",
			unode:data.unode,
		}
		storageChannel<- sItem
	}

}
func uvCounter(uvChannel chan urlData, storageChannel chan storageBlock) {
	for data := range uvChannel {

		// HyperLoglog redis 去重
		hypperLoglogKey := "uv_hpll_"+getTime(data.data.time, "day")

		ret , err := redisClient.PFAdd(hypperLoglogKey,data.uid).Result()
		redisClient.Expire(hypperLoglogKey,time.Second*86400)
		if err !=nil {
			log.Fatalf("uvCounter PFAdd error=",err)
		}
		// 说明已经统计了
		if ret !=1 {
			continue
		}
		sItem := storageBlock{
			counterType:"uv",
			storageModel:"ZINCRBY",
			unode:data.unode,
		}
		storageChannel<- sItem
	}
}

func getTime(logTime,timeType string)string  {
	var item string
	switch timeType {
	case "day":
		item = "2006-01-02"
		break
	case "hour":
		item = "2006-01-02 15"
		break
	case "min":
		item = "2006-01-02 15:04"
		break
	}
	t, _ := time.Parse(item, time.Now().Format(item))

	return strconv.FormatInt(t.Unix(),10)
}

func logConsumer(logChannel chan string, pvChannel, uvChannel chan urlData) error {
	for  logStr := range logChannel{
		//切割日志字符串，扣除打点上报的数据
		data := cutLogFetchData(logStr)

		// uid 先用 refer+ua md5生成
		hash := md5.New()
		hash.Write([]byte(data.ua+data.refer))
		uid := hex.EncodeToString(hash.Sum(nil))

		uData := urlData{
			data:data,
			uid:uid,
			unode: formatUrl(data.url,data.time),
		}
		//log.Infof("logConsumer %v",uData)
		pvChannel <- uData
		uvChannel <- uData
	}
	return nil
}

func formatUrl(url,time string) urlNode  {
	regexpRule := regexp.MustCompile(`/([\d]+)/(([\d]+)|([\d]?))`)
	regexpMatch := regexpRule.FindStringSubmatch(url)
	if len(regexpMatch) == 0 {
		return urlNode{
			unType:"index",
			unUrl:url,
			unRequestId:"0",
			unTime:time,
		}
	} else {
		if len(regexpMatch[3]) >0 {
			return urlNode{
				unType:"detail",
				unUrl:url,
				unRequestId:regexpMatch[0],
				unTime:time,
			}
		} else {
			return urlNode{
				unType:"list",
				unUrl:url,
				unRequestId:regexpMatch[0],
				unTime:time,
			}
		}
	}
}

func cutLogFetchData(logStr string)digData {
	logStr = strings.TrimSpace(logStr)

	// 字符串截取,截取出所需内容
	pos1 := str.IndexOf(logStr,HANDLE_GIF,0)
	if pos1 == -1 {
		return digData{}
	}
	pos1 += len(HANDLE_GIF)
	pos2 := str.IndexOf(logStr," HTTP/", pos1)

	// 从pos1 的位置开始截取，截取出pos2-pos1 个长度的logStr
	urlParams := str.Substr(logStr, pos1, pos2-pos1)

	// 解析url参数，由于url.Parse()只认完整的网站因此，需要拼上协议及域名
	urlInfo, err := url.Parse("http://localhost:80/?"+urlParams)
	if err !=nil {
		return digData{}
	}

	data := urlInfo.Query()

	return digData{
		time:data.Get("time"),
		url:data.Get("url"),
		refer:data.Get("refer"),
		ua:data.Get("ua"),
	}
}

func readFileLineByLine(params cmdParams, logChannel chan string) error{
	fd, err := os.Open(params.logFilePath)
	if err != nil {
		log.Warningf("readFileLineByLine os.Open(%s) error =%s",params.logFilePath,err.Error())
		return err
	}
	defer fd.Close()
	bufRead := bufio.NewReader(fd)
	count := 0
	for {
		line, err := bufRead.ReadString('\n')
		//log.Infof("line:",line)
		if err != nil {
			if err == io.EOF {
				time.Sleep(time.Second * 3)
				log.Infof("readFileLineByLine wait readLine :%d", count)
			} else {
				log.Warningf("readFileLineByLine bufRead.ReadString error = %s",err.Error())
			}
		}
		logChannel <- line
		count++

		if count % (params.routineNum*1000) == 0 {
			log.Infof("readFileLineByLine line:%d", count)
		}
	}

	return nil
}
