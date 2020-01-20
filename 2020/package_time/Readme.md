# Go语言基础之time包

[源自:李文周老师的博客]: https://www.liwenzhou.com/posts/Go/go_time/

## time包

time包提供了时间的显示和测量用的函数。日历的计算采用的是公历。

### 1. 时间类型

`time.Time`类型表示时间。我们可以通过`time.Now()`函数获取当前的时间对象，然后获取时间对象的年月日时分秒等信息。示例代码如下：

```go
func timeDemo() {
	// 获取当前时间
	now := time.Now()
	fmt.Printf("当前时间:%v\n",now)
	year := now.Year()
	fmt.Printf("年:%d\n",year)
	month := now.Month()
	fmt.Printf("月:%02d\n",month)
	day := now.Day()
	fmt.Printf("日:%02d\n",day)
	hour := now.Hour()
	fmt.Printf("时:%2d\n", hour)
	min := now.Minute()
	fmt.Printf("分:%2d\n", min)
	second := now.Second()
	fmt.Printf("秒:%02d\n", second)
	currentTimeStr := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, min, second)
	fmt.Printf("当前时间:%s\n",currentTimeStr)
}
```

输出结果:

```go
当前时间:2020-01-08 23:33:23.668895 +0800 CST m=+0.000204850
年:2020
月:01
日:08
时:23
分:33
秒:23
当前时间:2020-01-08 23:33:23
```



### 2.时间戳

时间戳是自1970年1月1日（08:00:00GMT）至当前时间的总毫秒数。它也被称为Unix时间戳（UnixTimestamp）。

基于时间对象获取时间戳的示例代码如下：

```go
func timestampDemo() {
	// 获取当前时间
	now := time.Now()
	// 时间戳.
	timestamp1 := now.Unix()
	// 纳秒时间戳
	timestamp2 := now.UnixNano()
	// 获取xxxx-xx-xx-xx xx:xx:xx.xxx 类型时间
	currentTime := now.Format("2006-01-02 15:04:05.000")
	fmt.Printf("当前时间:%s\n",currentTime)
	fmt.Printf("时间戳%d\n", timestamp1)
	fmt.Printf("纳秒时间戳%d\n", timestamp2)
}
```

输出结果:

```go
当前时间:2020-01-08 23:43:49.440
时间戳1578498229
纳秒时间戳1578498229440737000
```

使用`time.Unix()`函数可以将时间戳转为时间格式

```go
func timestamp2Demo() {
	//timestamp := time.Now().Unix()
	timestamp := 1535883680 // 2018/9/2 18:21:20
	// 将时间戳转为时间格式
	timeObj := time.Unix(int64(timestamp), 0)
	fmt.Printf("当前时间:%s\n",timeObj.Format("2006/01/02 15:04:05"))
	fmt.Printf("年:%d\n",timeObj.Year())
	fmt.Printf("月:%02d\n",timeObj.Month())
	fmt.Printf("日:%02d\n",timeObj.Day())
	fmt.Printf("时:%02d\n", timeObj.Hour())
	fmt.Printf("分:%02d\n", timeObj.Minute())
	fmt.Printf("秒:%02d\n", timeObj.Second())
}
```

输出结果:

```go
当前时间:2018/09/02 18:21:20
年:2018
月:09
日:02
时:18
分:21
秒:20
```

### 3.时间间隔

`time.Duration`是`time`包定义的一个类型，它代表两个时间点之间经过的时间，以纳秒为单位。`time.Duration`表示一段时间间隔，可表示的最长时间段大约290年。

time包中定义的时间间隔类型的常量如下：

```go
const (
    Nanosecond  Duration = 1
    Microsecond          = 1000 * Nanosecond
    Millisecond          = 1000 * Microsecond
    Second               = 1000 * Millisecond
    Minute               = 60 * Second
    Hour                 = 60 * Minute
)
```

例如：`time.Duration`表示1纳秒，`time.Second`表示1秒

### 4.时间操作

#### 1）add

我们在日常的编码过程中可能会遇到要求时间+时间间隔的需求，Go语言的时间对象有提供Add方法如下：

```go
func (t Time) Add(d Duration) Time
```

举个例子，求一个小时之后的时间：

```go
func testAdd() {
	now := time.Now()
	fmt.Printf("当前时间:%s\n",now.Format("2006/01/02 15:04:05"))
	plusOneHouTime := now.Add(time.Hour)
	fmt.Printf("加一小时后的时间:%s\n", plusOneHouTime.Format("2006/01/02 15:04:05"))
	fmt.Printf("加3599秒的时间:%s\n", now.Add(time.Second * 3599).Format("2006/01/02 15:04:05"))
}
```

输出：

```go
当前时间:2020/01/09 00:05:02
加一小时后的时间:2020/01/09 01:05:02
加3600秒的时间:2020/01/09 01:05:01
```

#### 2) .sub

求两个时间之间的差值：

```go
func (t Time) Sub(u Time) Duration
```

返回一个时间段t-u。如果结果超出了Duration可以表示的最大值/最小值，将返回最大值/最小值。要获取时间点t-d（d为Duration），可以使用t.Add(-d)

```go
func testSub() {
  // 当前时间与当前时间两小时前的差值
	sub := time.Now().Sub(time.Now().Add(-2 * time.Hour))
	fmt.Println(sub)
}
```

输出:

```go
1h59m59.999999905s // 1小时59分59.99999905秒
```

#### 3).Equal

```go
func (t Time) Equal(u Time) bool
```

判断两个时间是否相同，会考虑时区的影响，因此不同时区标准的时间也可以正确比较。本方法和用t==u不同，这种方法还会比较地点和时区信息。

代码:

```go
func testEqual() {
	now := time.Now()
	time1 := now.Add(time.Hour)
	if time1.Equal(now.Add(3600 * time.Second)) {
		fmt.Println("加一小时的时间与加3600秒的时间相等")
	} else {
		fmt.Println("加一小时的时间与加3600秒的时间不相等")
	}
}
```

输出:

```go
加一小时的时间与加3600秒的时间相等
```

#### 4).Before

```go
func (t Time) Before(u Time) bool
```

如果t代表的时间点在u之前，返回真；否则返回假。

5).After

```go
func (t Time) After(u Time) bool
```

如果t代表的时间点在u之后，返回真；否则返回假。

### 定时器

使用`time.Tick(时间间隔)`来设置定时器，定时器的本质上是一个通道（channel）。

```go
func tickDemo() {
	ticker := time.Tick(2 * time.Second)
	count := 0
	for i := range ticker {
		count++
		if count >=5 {
			break
		}
		fmt.Println(i.Format("2006/01/02 15:04:05"))
	}
}
```

输出:

```go
2020/01/09 00:36:00
2020/01/09 00:36:02
2020/01/09 00:36:04
2020/01/09 00:36:06
```

### 时间格式化

时间类型有一个自带的方法`Format`进行格式化，需要注意的是Go语言中格式化时间模板不是常见的`Y-m-d H:M:S`而是使用Go的诞生时间2006年1月2号15点04分（记忆口诀为2006 1 2 3 4）。也许这就是技术人员的浪漫吧。

补充：如果想格式化为12小时方式，需指定`PM`

```go
func formatDemo() {
	now := time.Now()
	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
}
```

### 解析字符串格式的时间

```go
now := time.Now()
fmt.Println(now)
// 加载时区
loc, err := time.LoadLocation("Asia/Shanghai")
if err != nil {
	fmt.Println(err)
	return
}
// 按照指定时区和指定格式解析字符串时间
timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2019/08/04 14:15:20", loc)
if err != nil {
	fmt.Println(err)
	return
}
fmt.Println(timeObj)
fmt.Println(timeObj.Sub(now))
```