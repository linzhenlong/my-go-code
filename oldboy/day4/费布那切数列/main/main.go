package main

import "fmt"

func main()  {

	var a [100]uint64

	a[0] = 1;
	a[1] = 1;
	for i:=2;i<100;i++ {
		a[i] = a[i-1] + a[i-2]
	}

	for key,val:=range a {
		fmt.Printf("a[%d]=%d\n",key,val)
	}
}
