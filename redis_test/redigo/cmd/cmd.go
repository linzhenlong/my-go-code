package cmd

import (
	"fmt"
	"github.com/linzhenlong/my-go-code/redis_test/redigo/pool"
)

func Get()  {
	redis := pool.New()
	fmt.Printf("%v", redis)

}
