package main

import (
	"fmt"
	"math/rand"
)

type job struct {
	ID      int
	RandNum int
}

type result struct {
	Job *job
	sum int
}

// 创建工作池
func createPool(workerNum int, jobChan chan *job, resChan chan *result) {
	for i := 0; i < workerNum; i++ {
		// 匿名函数闭包
		go func() {
			// 执行运算
			for job := range jobChan {
				randNum := job.RandNum
				// 计算结果，随机数的每一位求和
				var sum int
				for randNum != 0 {
					tmp := randNum % 10
					sum += tmp
					randNum /= 10
				}
				res := &result{
					Job: job,
					sum: sum,
				}
				resChan <- res

			}
		}()
	}
}

func main() {
	// 2个管道
	// 1.job
	jobChan := make(chan *job, 32)

	// 2.结果管道
	resChan := make(chan *result, 32)

	workerNum := 5
	createPool(workerNum, jobChan, resChan)

	// 打印协程
	go func() {
		//遍历结果集打印
		for res := range resChan {
			fmt.Printf("job_ID:%d,randNum:%d,sum:%d\n", res.Job.ID, res.Job.RandNum, res.sum)
		}
	}()
	var id int
	for {
		id++
		//rand.Seed(time.Now().UnixNano())
		randNum := rand.Int()
		job := &job{
			ID:      id,
			RandNum: randNum,
		}
		jobChan <- job
	}

}
