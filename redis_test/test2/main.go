package main

import "fmt"

func main()  {
	c();
}
func c()  {
	fmt.Println("redis")
}

/*func c()*redis.Client  {
	client := redis.NewClient(&redis.Options{
		Addr:"127.0.0.1",
		Password:"",
		DB:0,
	})
	fmt.Printf("%v",client);
}*/

