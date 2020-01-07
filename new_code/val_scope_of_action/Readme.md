# Golang变量的作用域问题

1. `函数`内部声明/定义的变量叫做`局部变量`，作用域`仅限于函数内部`

```go
package main

import "fmt"

func test() {
	name := "tom"
	age :=19
	fmt.Println("func test", name,age)
}

func main() {
	fmt.Println(name) // 会报错，因为name 是局部变量
}
```

2. `函数外`部声明/定义的变量叫`全局变量`，作用域在整个包都有效，如果其首字母为大写，则作用域在整个程序有效

```go
package main

import "fmt"

var name string = "张三"
func test() {
	name := "tom"
	age :=19
	fmt.Println("func test", name,age)
}

func main() {
	test()
	fmt.Println(name) // 输出为张三，name虽然定义了全局变量，并且在test函数中赋值了，但是，test 中赋值作用域为test函数
}
```



3. 如果变量是在一个代码块，比如`for/if`中，那么这个变量的作用域就在这个代码块内

例子1：

```go
package main

import "fmt"

var name string = "张三"
func test() {
	name := "tom"
	age :=19
	fmt.Println("func test", name,age)
}

func main() {
	test()
	fmt.Println(name) // 输出为张三，name虽然定义了全局变量，并且在test函数中赋值了，但是，test 中赋值作用域为test函数

	i := 18
	for i:=0;i<5;i++ {
		fmt.Println("i",i)
	}
	fmt.Println("i:=8的i==>",i)
}
```

输出:

```go
~/webroot/golang/my-go-code/new_code/val_scope_of_action ⮀ ⭠ develop± ⮀ go run main.go
func test tom 19
张三
i 0
i 1
i 2
i 3
i 4
i:=8的i==> 18
```



例子2：

```go

package main

import "fmt"

var name string = "张三"
func test() {
	name := "tom"
	age :=19
	fmt.Println("func test", name,age)
}

func main() {
	test()
	fmt.Println(name) // 输出为张三，name虽然定义了全局变量，并且在test函数中赋值了，但是，test 中赋值作用域为test函数

	i := 18
	for i=0;i<5;i++ {
		fmt.Println("i",i)
	}
	fmt.Println("i:=18的i==>",i)
}
```

输出:

```go
 ~/webroot/golang/my-go-code/new_code/val_scope_of_action ⮀ ⭠ develop± ⮀ go run main.go
func test tom 19
张三
i 0
i 1
i 2
i 3
i 4
i:=8的i==> 5 // 相当于对代码块外部的变量进行重新赋值
```

