# Golang的context标准库

> 源自:[李文周老师的博客](https://www.liwenzhou.com/posts/Go/go_context/)

在 Go http包的Server中，每一个请求在都有一个对应的 goroutine 去处理。

请求处理函数通常会启动额外的 goroutine 用来访问后端服务，比如数据库和RPC服务。

用来处理一个请求的 goroutine 通常需要访问一些与请求特定的数据，比如终端用户的身份认证信息、验证相关的token、请求的截止时间。 

当一个请求被取消或超时时，所有用来处理该请求的 goroutine 都应该迅速退出，然后系统才能释放这些 goroutine 占用的资源



## 0.为什么需要Context?

Sync.WaitGroup 简单示例

```GO
package main

import "sync"
import "net/http"
import "fmt"

//https://go-zh.org/pkg/sync/#example_WaitGroup

var wg sync.WaitGroup

var urls = []string{
	"http://www.baidu.com",
	"http://www.qq.com",
	"http://www.abc.com",
	"http://www.tmall.com",
}

func main() {
	for _, url := range urls {
		// 递增WaitGroup 计数器
		wg.Add(1)
		// 启动一个goroutine来取回URL
		go func(url string) {
			// 当goroutine 执行完成，减小计数器
			defer wg.Done()

			response,err := http.Get(url)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(url,response.StatusCode)
		}(url)
	}
	// 等待所有的HTTP取回操作.
	wg.Wait()
}
```

基本示例

```go
package main

import "sync"
import "fmt"
import "time"

var wg sync.WaitGroup

func Worker() {
	for {
		fmt.Println("worker")
		time.Sleep(time.Second * 1)
	}

	// 如何接收外部命令退出
	wg.Done()
}

func main() {
	wg.Add(1)
	go Worker()
	// 如果优雅的结束子goroutine
	wg.Wait()
	fmt.Println("over")
}
```



全局变量的方式

```go
package main

import "sync"
import "fmt"
import "time"

var wg sync.WaitGroup

// 全局变量控制是否退出
var exit bool

// 全局变量方式存在的问题
// 1.使用全局变量在跨包调用时不容易统一
// 2.如果Process中再启动goroutine，就不太好控制了.

func Process() {
	for {
		fmt.Println("process run...")
		time.Sleep(time.Second)
		if exit {
			break
		}
	}
	wg.Done()
}

func main() {
	wg.Add(1)
	go Process()
	// sleep 3秒以免程序过快退出
	time.Sleep(time.Second * 3)
	// 修改全局变量实现子goroutine的退出
	exit = true
	wg.Wait()
	fmt.Println("main over")

}

```

管道(channel)方式

```go
package main

import "sync"

import "fmt"

import "time"

var wg sync.WaitGroup

// 管道方式存在的问题：
// 1. 使用全局变量在跨包调用时不容易实现规范和统一，需要维护一个共用的channel

func Process(exitChan chan struct{}) {
	LOOP:
	for {
		fmt.Println("process run...")
		time.Sleep(time.Second)
		select {
		case <-exitChan: // 等待接收上级通知
		break LOOP
		default:
		}
	}
	wg.Done()
}

func main() {
	var exitChan = make(chan struct{})
	wg.Add(1)
	go Process(exitChan)
	// sleep 3秒以免程序过快退出
	time.Sleep(time.Second*3)

	// 给子goroutine一个退出的信号
	exitChan<- struct{}{} 
	close(exitChan)
	wg.Wait()
	fmt.Println("main over...")
}

```

官方版的方案

```go
package main

import "sync"

import "context"

import "time"

import "fmt"


var wg sync.WaitGroup

func Process(ctx context.Context) {
	LOOP:
	for {
		fmt.Println("process run...")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():// 等待上级通知
		break LOOP
		default:
		}
	}
	wg.Done()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go Process(ctx)
	time.Sleep(time.Second * 4)

	// 通知子goroutine结束
	cancel() 
	wg.Wait()
	fmt.Println("main over")
} 
```

当子goroutine又开启另外一个goroutine时，只需要将ctx传入即可：

```go
package main

import "sync"

import "context"

import "fmt"

import "time"

var wg sync.WaitGroup

func Process2(ctx context.Context) {
	LOOP:
	for {
		fmt.Println("Process2 run...")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
}

func Process(ctx context.Context) {
	go Process2(ctx)
	LOOP:
	for {
		fmt.Println("Process run...")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
	wg.Done()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go Process(ctx)
	time.Sleep(time.Second * 3)
	cancel()
	wg.Wait()
	fmt.Println("main over...")

}
```



## 1. Context初识

Go1.7加入了一个新的标准库context，它定义了Context类型,专门用来简化,对于处理单个请求的多个 goroutine 之间与请求域的数据、取消信号、截止时间等相关操作，这些操作可能涉及多个 API 调用。

对服务器传入的请求应该创建上下文，而对服务器的传出调用应该接受上下文。它们之间的函数调用链必须传递上下文，或者可以使用WithCancel、WithDeadline、WithTimeout或WithValue创建的派生上下文。当一个上下文被取消时，它派生的所有上下文也被取消。



## 2. Context接口

`context.Context` 是一个接口，该接口定义了四个需要实现的方法。具体签名如下：

```go
type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}
```

其中：

- Deadline方法需要返回当前Context被取消的时间，也就是完成工作的截止时间（deadline）；
- Done方法需要返回一个Channel，这个Channel会在当前工作完成或者上下文被取消之后关闭，多次调用Done方法会返回同一个Channel；
- Err方法会返回当前Context结束的原因，它只会在Done返回的Channel被关闭时才会返回非空的值；
  - 如果当前Context被取消就会返回Canceled错误；
  - 如果当前Context超时就会返回DeadlineExceeded错误；
- Value方法会从Context中返回键对应的值，对于同一个上下文来说，多次调用Value 并传入相同的Key会返回相同的结果，该方法仅用于传递跨API和进程间跟请求域的数据；

## 3. Background()和TODO()

Go内置两个函数：Background()和TODO()，这两个函数分别返回一个实现了Context接口的background和todo。我们代码中最开始都是以这两个内置的上下文对象作为最顶层的partent context，衍生出更多的子上下文对象。

Background()主要用于main函数、初始化以及测试代码中，作为Context这个树结构的最顶层的Context，也就是根Context。

TODO()，它目前还不知道具体的使用场景，如果我们不知道该使用什么Context的时候，可以使用这个。

background和todo本质上都是emptyCtx结构体类型，是一个不可取消，没有设置截止时间，没有携带任何值的Context。



## 4.With系列函数

此外，`context`包中还定义了四个With系列函数。

### 4.1 WithCancel

WithCancel的函数签名如下：

```go
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
```

WithCancel返回带有新Done通道的父节点的副本。当调用返回的cancel函数或当关闭父上下文的Done通道时，将关闭返回上下文的Done通道，无论先发生什么情况。

取消此上下文将释放与其关联的资源，因此代码应该在此上下文中运行的操作完成后立即调用cancel。

```go
func gen(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // return结束该goroutine，防止泄露
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 当我们取完需要的整数后调用cancel

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}
```

上面的示例代码中，gen函数在单独的goroutine中生成整数并将它们发送到返回的通道。 gen的调用者在使用生成的整数之后需要取消上下文，以免gen启动的内部goroutine发生泄漏。

### 4.2 WithDeadline

`WithDeadline`的函数签名如下：

```go
func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
```

返回父上下文的副本，并将deadline调整为不迟于d。如果父上下文的deadline已经早于d，则WithDeadline(parent, d)在语义上等同于父上下文。当截止日过期时，当调用返回的cancel函数时，或者当父上下文的Done通道关闭时，返回上下文的Done通道将被关闭，以最先发生的情况为准。

取消此上下文将释放与其关联的资源，因此代码应该在此上下文中运行的操作完成后立即调用cancel。

```go
package main

import "time"
import "context"
import "fmt"

func main() {
	d := time.Now().Add(50 * time.Millisecond)
	//d := time.Now().Add(2 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// 尽管ctx会过期，但在任何情况下调用它的cancel函数都是很好的实践。
	// 如果不这样做，可能会使上下文及其父类存活的时间超过必要的时间。
	defer cancel()
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("over slept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
```

上面的代码中，定义了一个50毫秒之后过期的deadline，然后我们调用`context.WithDeadline(context.Background(), d)`得到一个上下文（ctx）和一个取消函数（cancel），然后使用一个select让主程序陷入等待：等待1秒后打印`overslept`退出或者等待ctx过期后退出。 因为ctx50毫秒后就过期，所以`ctx.Done()`会先接收到值，上面的代码会打印ctx.Err()取消原因。



### 4.3 WithTimeout

`WithTimeout`的函数签名如下：

```go
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
```

`WithTimeout`返回`WithDeadline(parent, time.Now().Add(timeout))`。

取消此上下文将释放与其相关的资源，因此代码应该在此上下文中运行的操作完成后立即调用cancel，通常用于数据库或者网络连接的超时控制。具体示例如下：

```go
package main

import "sync"
import "context"
import "time"
import "fmt"

var wg sync.WaitGroup

func Process(ctx context.Context) {
	LOOP:
	for {
		fmt.Println("db connecting....")
		time.Sleep(time.Millisecond * 10)
		select {
		case <-ctx.Done(): //50ms后自动调用
		break LOOP
		default:
		}
	}
	fmt.Println("Process over...")
	wg.Done()
}

func main() {
	// 50ms 超时.
	ctx, cancel := context.WithTimeout(context.Background(),time.Millisecond * 50 ) 
	wg.Add(1)
	go Process(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
	fmt.Println("main over")
}
```

### 4.4 WithValue

`WithValue`函数能够将请求作用域的数据与 Context 对象建立关系。声明如下：

```go
func WithValue(parent Context, key, val interface{}) Context
```

`WithValue`返回父节点的副本，其中与key关联的值为val。

仅对API和进程间传递请求域的数据使用上下文值，而不是使用它来传递可选参数给函数。

所提供的键必须是可比较的，并且不应该是`string`类型或任何其他内置类型，以避免使用上下文在包之间发生冲突。`WithValue`的用户应该为键定义自己的类型。为了避免在分配给interface{}时进行分配，上下文键通常具有具体类型`struct{}`。或者，导出的上下文关键变量的静态类型应该是指针或接口。



```go
package main

import "sync"

import "context"

import "time"

import "fmt"

type TraceCode string

var wg sync.WaitGroup


func Process2(ctx context.Context) {
	key1 := TraceCode("TRACE_CODE")
	traceCode1, ok := ctx.Value(key1).(string) // 在子goroutine中获取trace code

	key2 := TraceCode("TRACE_CODE2")
	traceCode2, ok := ctx.Value(key2).(string) // 在子goroutine中获取trace code

	if !ok {
		fmt.Println("非法的trace code...")
	}
	LABLE:
	for {
		fmt.Printf("Process2 trace_code1=>%s,trace_code12=>%s\n",traceCode1, traceCode2)
		time.Sleep(time.Millisecond * 5)
		select {
		case <-ctx.Done():
			break LABLE
		default:
		}
	}

}

func Process(ctx context.Context) {
	key := TraceCode("TRACE_CODE")
	traceCode, ok := ctx.Value(key).(string) // 在子goroutine中获取trace code

	if !ok {
		fmt.Println("非法的trace code...")
	}
	ctx = context.WithValue(ctx,TraceCode("TRACE_CODE2"), "77889966")
	go Process2(ctx)
	LABLE:
	for {
		fmt.Printf("Process, trace code:%s\n",traceCode)
		time.Sleep(time.Millisecond * 10)
		select {
		case <-ctx.Done():
			break LABLE
		default:

		}
	}
	fmt.Println("Process done....")
	wg.Done()
}

func main() {

	// 设置50ms的超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond * 50)

	// 在系统的入口设置trace code 传递给后续启动的goroutine实现日志聚合
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "12512312234")
	wg.Add(1)
	go Process(ctx)
	time.Sleep(time.Second * 5)
	cancel() // 通知子goroutine 结束
	wg.Wait()
	fmt.Println("main over...")

}
```



## 5. 使用Context的注意事项

- 推荐以参数的方式显示传递Context
- 以Context作为参数的函数方法，应该把Context作为第一个参数。
- 给一个函数方法传递Context的时候，不要传递nil，如果不知道传递什么，就使用context.TODO()
- Context的Value相关方法应该传递请求域的必要数据，不应该用于传递可选参数
- Context是线程安全的，可以放心的在多个goroutine中传递



## 6.客户端超时取消示例

调用服务端API时如何在客户端实现超时控制？

### server端

```go
package main

import "net/http"

import "math/rand"

import "time"

import "fmt"

import "strconv"

func indexHandler(w http.ResponseWriter,r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(2)
	if randNum == 0 {
		time.Sleep(time.Second * 3) // 模拟一下服务端超时
		fmt.Fprintf(w,"slow response"+strconv.Itoa(randNum))
		return
	}
	fmt.Fprintf(w,"quick response"+strconv.Itoa(randNum))
}

func main() {
	http.HandleFunc("/",indexHandler)
	err := http.ListenAndServe("0.0.0.0:8889",nil)
	if err !=nil {
		panic(err)
	}
}
```

### client端

```go
package main

import "context"

import "time"

import "net/http"

import "fmt"

import "sync"

import "io/ioutil"


type respData struct {
	resp *http.Response
	err error
}

func doCall(ctx context.Context) {
	transport := http.Transport{
		// 请求频繁可定义全局的Client对象并开启长连接
		// 请求不频繁使用短连接
		DisableKeepAlives: true,
	}
	client := http.Client{
		Transport: &transport,
	}
	respChan := make(chan *respData,1)
	req,err := http.NewRequest("GET","http://127.0.0.1:8889", nil)
	if err != nil {
		fmt.Printf("new requestg failed, err:%v\n", err)
		return
	}
	req = req.WithContext(ctx) // 使用带超时的ctx创建一个新的client request

	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Wait()
	go func(){
		resp , err := client.Do(req)
		fmt.Printf("client DO resp:%v,err:%v\n",resp, err)
		rd := &respData{
			resp: resp,
			err: err,
		}
		respChan<- rd
		wg.Done()
	}()
	select{
	case <-ctx.Done():
		fmt.Println("call api timeout")
	case result := <- respChan:
		fmt.Println("call api succ")
		if result.err != nil {
			fmt.Printf("call server api failed, err:%v\n", result.err)
			return
		}
		defer result.resp.Body.Close()
		data,_ := ioutil.ReadAll(result.resp.Body)
		fmt.Printf("resp:%v\n",string(data))
	}
}

func main() {
	// 定义一个100ms超时的上下文
	ctx,cancel := context.WithTimeout(context.Background(),time.Millisecond * 100)

	defer cancel()

	// 请求服务端.
	doCall(ctx)
}
```

