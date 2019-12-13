package cmd

import (
	"fmt"
	"go_dev/redis_test/redigo/pool"
)

func Get()  {
	redis := pool.New()
	fmt.Printf("%v", redis)

}
