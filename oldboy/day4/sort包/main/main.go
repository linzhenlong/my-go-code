package main

import (
	"fmt"
	"sort"
)

func main()  {
	testIntSort()
	testStringSort()
	testFloatSort()
	testIntSearch()
	testFloatSearch()
	testStringSearch()
}

func testIntSort()  {
	var intArr = [...]int{1,3,3434,55,677,99,2,10}
	sort.Ints(intArr[:])
	fmt.Println(intArr)
}

func testStringSort() {
	var strArr = [...]string{"azc","bac","bz","a"}
	sort.Strings(strArr[:])
	fmt.Println(strArr)
}

func testFloatSort()  {
	var a = [...]float64{0.3,0.7,2.9,4.5,0.001}
	sort.Float64s(a[:])
	fmt.Println(a)
}

func testIntSearch() {
	var intArr = [...]int{1,3,3434,55,677,99,2,10}
	sort.Ints(intArr[:])
	index := sort.SearchInts(intArr[:],2)
	fmt.Println(intArr[index])
}

func testFloatSearch()  {
	var a = [...]float64{0.3,0.7,2.9,4.5,0.001}
	sort.Float64s(a[:])
	index := sort.SearchFloat64s(a[:],2.9)
	fmt.Println(a[index])
}

func testStringSearch()  {
	var strArr = [...]string{"azc","bac","bz","a"}
	sort.Strings(strArr[:])
	index := sort.SearchStrings(strArr[:],"a")
	fmt.Println(index)
}