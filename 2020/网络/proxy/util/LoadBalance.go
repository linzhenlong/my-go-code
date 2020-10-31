package util

import (
	"fmt"
	"hash/crc32"
	"log"
	"math/rand"
	"sort"
	"time"
)

var (
	// LB 对象.
	LB *LoadBalance
	// ServerIndices web server 权重数组，
	// 假设两台web服务器序号是0，1,权重是2:4 那么数组里就放2个0，4个1[0,0,1,1,1,1]
	ServerIndices []int
	// SumWeight 总权重
	SumWeight int
)

// HTTPServers ...
type HTTPServers []*HTTPServer

// 实现sort包的sort接口.
func (p HTTPServers) Len() int {
	return len(p)
}
func (p HTTPServers) Less(i, j int) bool {
	return p[i].CurrentWeight > p[j].CurrentWeight // 从大到小.
}
func (p HTTPServers) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// HTTPServer 自定义http server
type HTTPServer struct {
	Host          string // http://localhost:9090或 http://localhost:9091
	Weight        int    // 权重
	CurrentWeight int    // 当前权重 (平滑加权算法会用到)
	Status        string // web服务器状态,默认UP,宕机了使用DOWN
}

// LoadBalance 负载均衡...
type LoadBalance struct {
	Servers      HTTPServers
	CurrentIndex int // 指向当前的访问的web服务器.
}

// NewLoadBalance 构造函数
func NewLoadBalance() *LoadBalance {
	return &LoadBalance{Servers: make([]*HTTPServer, 0)}
}

// NewHTTPServer 构造函数
func NewHTTPServer(host string, weight int) *HTTPServer {
	return &HTTPServer{
		Host:          host,
		Weight:        weight,
		CurrentWeight: 0, // 初始权重为0
		Status:        "UP",
	}
}

func init() {
	// 负载均衡
	LB = NewLoadBalance()
	// 先写死

	webServer1 := NewHTTPServer("http://localhost:9091", 3)
	webServer2 := NewHTTPServer("http://localhost:9092", 1)
	webServer3 := NewHTTPServer("http://localhost:9093", 1)
	LB.AddServer(webServer1)
	LB.AddServer(webServer2)
	LB.AddServer(webServer3)

	for index, server := range LB.Servers {
		if server.Weight > 0 {
			for i := 0; i < server.Weight; i++ {
				ServerIndices = append(ServerIndices, index)
			}
		}
		// 总权重..
		SumWeight = SumWeight + server.Weight
	}
	// 检测web服务器状态,切协程跑
	go func() {
		checkServers(LB.Servers)
	}()

	log.Printf("ServerIndices:%v", ServerIndices)
}

// AddServer 添加server.
func (l *LoadBalance) AddServer(server *HTTPServer) {
	l.Servers = append(l.Servers, server)
}

// SelectByRand 随机算法
func (l *LoadBalance) SelectByRand() *HTTPServer {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(l.Servers))
	return l.Servers[index]
}

// SelectByIPHash ip_hash 算法
func (l *LoadBalance) SelectByIPHash(ip string) *HTTPServer {
	n := crc32.ChecksumIEEE([]byte(ip))
	index := n % uint32(len(l.Servers))
	log.Printf("ip:%s,hash_code:%d\n", ip, n)
	return l.Servers[index]
}

// SelectByRandWithWeight 随机算法
func (l *LoadBalance) SelectByRandWithWeight() *HTTPServer {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(len(ServerIndices))
	index := ServerIndices[randNum]
	return l.Servers[index]
}

// SelectByRandWithWeight2 随机算法改良升级版
// 举个例子
// 假设A:B:C的比重是5:2:1
// 我们针对这三个数进行变量并计算权重总和,变成了5,7(5+2),8(5+2+1)
// 这样我们心里就形成了3个区间[0,5),[5,7),[7,8)
// 然后在[0,8)之内取一个随机数,落到那个区间,就是哪个.
func (l *LoadBalance) SelectByRandWithWeight2() *HTTPServer {
	rand.Seed(time.Now().UnixNano())
	sumList := make([]int, len(l.Servers))
	sum := 0
	for i := 0; i < len(l.Servers); i++ {
		sum += l.Servers[i].Weight
		sumList[i] = sum
	}
	rad := rand.Intn(sum)
	for index, value := range sumList {
		if rad < value {
			return l.Servers[index]
		}
	}
	return l.Servers[0]
}

// RoundRobin 轮询算法
func (l *LoadBalance) RoundRobin() *HTTPServer {
	curServer := l.Servers[l.CurrentIndex]
	// 进行轮询，判断同时判断一下别越界
	l.CurrentIndex++
	if l.CurrentIndex >= len(l.Servers) {
		l.CurrentIndex = 0
	}
	return curServer
}

// RoundRobinWithWeight 加权轮询算法
// 例如，三台服务器,比重3:1:1= 5
// 按照我们加权随机算法的思路,可以这样操作
// 初始化一个数组
// [0,0,0,1,2]
// [s1,s1,s1,s2,s3]
// 然后按照轮询算法一个个来
// 缺点：
// 很明显，会造成s1的压力过大。我们更希望的5次请求内是[s1,s2,s1,s3,s1],这就叫平滑加权
func (l *LoadBalance) RoundRobinWithWeight() *HTTPServer {
	curServer := l.Servers[ServerIndices[l.CurrentIndex]]
	// 进行加权轮询
	l.CurrentIndex = (l.CurrentIndex + 1) % len(ServerIndices)
	log.Println("CurrentIndex:", l.CurrentIndex)
	return curServer
}

// RoundRobinWithWeight2 加权随机，例如区间的方式
// 比如比重是3:1:1
// [0,3),[3,4),[4,5)
func (l *LoadBalance) RoundRobinWithWeight2() *HTTPServer {
	server := l.Servers[l.CurrentIndex]
	sum := 0
	for i := 0; i < len(l.Servers); i++ {
		sum += l.Servers[i].Weight
		if l.CurrentIndex < sum {
			server = l.Servers[i]
			if l.CurrentIndex == sum-1 && i != len(l.Servers)-1 {
				l.CurrentIndex++
			} else {
				l.CurrentIndex = (l.CurrentIndex + 1) % sum
			}
			log.Println(l.CurrentIndex)
			break
		}
	}
	return server
}

// RoundRobinWithWeight3 平滑加权轮询算法
// 1.初始权重{s1:3,s2:1,s3:1},注意总权重是5
// 2.每次命中权重最大的返回,然后把命中节点的当前权重减去总权重
// 3.第二返回前,把当前权重加上原始权重
// 这样的做法是为了保证在5次（总权重）内，能恢复到格格权重是0

// RoundRobinWithWeight3 ... 一个循环内
/*
权重                                 命中                        命中后的权重
{s1:3,s2:1,s3:1} 初始权重            s1(因为他最大)        {s1:-2,s2:1,s3:1} s1要减去5
{s1:1,s2:2,s3:2} s1要加3，其他加1     s2                  {s1:1,s2:-3,s3:2} s2要减5
{s1:4,s2:-2,s3:3} s1要加3，其他加1    s1                  {s1:-1,s2:-2,s3:3} s1要减5
{s1:2,s2:-1,s3:4} s1要加3，其他加1    s3                  {s1:2,s2:-1,s3:-1} s3要减5
{s1:5,s2:0,s3:0} s1要加3，其他加1    s1                  {s1:0,s2:-1,s3:-1} s1要减5
*/
//
func (l *LoadBalance) RoundRobinWithWeight3() *HTTPServer {
	for _, s := range l.Servers {
		s.CurrentWeight = s.CurrentWeight + s.Weight
	}
	// 取出权重最大的
	sort.Sort(l.Servers)
	maxWeightServer := l.Servers[0]
	// 每次命中权重最大的返回,然后把命中节点的当前权重减去总权重
	maxWeightServer.CurrentWeight = maxWeightServer.CurrentWeight - SumWeight

	test := ""
	for _, s := range l.Servers {
		test += fmt.Sprintf("%d,", s.CurrentWeight)
	}
	fmt.Println(test)
	return maxWeightServer
}

func checkServers(servers HTTPServers) {
	ticker := time.NewTicker(time.Second * 3)
	cheker := NewHTTPChecker(servers)
	for {
		select {
		case <-ticker.C:
			cheker.Check(time.Second * 2)
			for _, s := range servers {
				log.Println("checkServers", s.Host, s.Status)
			}
			log.Println("======================")
		}
	}
}
