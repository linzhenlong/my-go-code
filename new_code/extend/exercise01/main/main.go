package main
import "fmt"
type A struct {
	Name string
	Age int
}
type Stu struct {
	A
	int
}
func main()  {
	stu := Stu{}
	stu.Name = "Tom"
	stu.Age = 10
	stu.int = 100
	fmt.Println(stu)
	fmt.Println(stu.int)
}
