package main

import "time"

import "fmt"

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

func testAdd() {
	now := time.Now()
	fmt.Printf("当前时间:%s\n",now.Format("2006/01/02 15:04:05"))
	plusOneHouTime := now.Add(time.Hour)
	fmt.Printf("加一小时后的时间:%s\n", plusOneHouTime.Format("2006/01/02 15:04:05"))
	fmt.Printf("加3599秒的时间:%s\n", now.Add(time.Second * 3599).Format("2006/01/02 15:04:05"))
}

func testSub() {
	sub := time.Now().Sub(time.Now().Add(-2 * time.Hour))
	fmt.Println(sub.Round(time.Second))
}

func testEqual() {
	now := time.Now()
	time1 := now.Add(time.Hour)
	if time1.Equal(now.Add(3600 * time.Second)) {
		fmt.Println("加一小时的时间与加3600秒的时间相等")
	} else {
		fmt.Println("加一小时的时间与加3600秒的时间不相等")
	}
}

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
func main() {
	timeDemo()
	fmt.Println("###timestampDemo###")
	timestampDemo()
	fmt.Println("###timestamp2Demo###")
	timestamp2Demo()

	testAdd()
	testSub()
	testEqual()
	tickDemo()
}