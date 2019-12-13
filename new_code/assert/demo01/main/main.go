package main
import "fmt"
type Point struct {
	X int
	Y int
}
func main()  {
	var a interface{}
	var point = Point{X:1,Y:2}
	a = point
	fmt.Println(a)
	var b Point
	// b = a  把 a直接赋给b 会报错
	// cannot use a (type interface {}) as type Point in assignment: need type assertion
	// 通过类型断言就可以了
	b = a.(Point)
	fmt.Println(b)
}
