package main

import "fmt"


type testRes struct {
	ID int
}
func test()(res *testRes) {
	return
}
func test2()(res testRes) {
	return
}
func test3()(res *testRes) {
	//res = new(testRes)
	res = &testRes{}
	return
}

func main() {
	r := test()
	fmt.Println(r)
	r2 := test2()
	fmt.Println(r2)
	r3 := test3()
	fmt.Println(r3,r3.ID)
}