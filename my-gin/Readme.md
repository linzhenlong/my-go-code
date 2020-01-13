# Go语言标准库之http/template

[源自:李文周老师的博客]:  https://www.liwenzhou.com/posts/Go/go_template/

`html/template`包实现了数据驱动的模板，用于生成可防止代码注入的安全的HTML内容。它提供了和`text/template`包相同的接口，Go语言中输出HTML的场景都应使用`html/template`这个包。

# 模板与渲染

在一些前后端不分离的Web架构中，我们通常需要在后端将一些数据渲染到HTML文档中，从而实现动态的网页（网页的布局和样式大致一样，但展示的内容并不一样）效果。

我们这里说的模板可以理解为事先定义好的HTML文档文件，模板渲染的作用机制可以简单理解为文本替换操作–使用相应的数据去替换HTML文档中事先准备好的标记。

很多编程语言的Web框架中都使用各种模板引擎，比如Python语言中Flask框架中使用的jinja2模板引擎。

### Go语言的模板引擎

Go语言内置了文本模板引擎`text/template`和用于HTML文档的`html/template`。它们的作用机制可以简单归纳如下：

1. 模板文件通常定义为`.tmpl`和`.tpl`为后缀（也可以使用其他的后缀），必须使用`UTF8`编码。
2. 模板文件中使用`{{`和`}}`包裹和标识需要传入的数据。
3. 传给模板这样的数据就可以通过点号（`.`）来访问，如果数据是复杂类型的数据，可以通过{ { .FieldName }}来访问它的字段。
4. 除`{{`和`}}`包裹的内容外，其他内容均不做修改原样输出。

### 模板引擎的使用

Go语言模板引擎的使用可以分为三部分：定义模板文件、解析模板文件和模板渲染.

### 定义模板文件

其中，定义模板文件时需要我们按照相关语法规则去编写

### 解析模板文件

上面定义好了模板文件之后，可以使用下面的常用方法去解析模板文件，得到模板对象：

```go
func (t *Template) Parse(src string) (*Template, error)
func ParseFiles(filenames ...string) (*Template, error)
func ParseGlob(pattern string) (*Template, error)
```

当然，你也可以使用`func New(name string) *Template`函数创建一个名为`name`的模板，然后对其调用上面的方法去解析模板字符串或模板文件。

### 模板渲染

渲染模板简单来说就是使用数据去填充模板，当然实际上可能会复杂很多。

```go
func (t *Template) Execute(wr io.Writer, data interface{}) error
func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error
```

### 基本示例

#### 定义模板文件

我们按照Go模板语法定义一个`hello.tmpl`的模板文件，内容如下：

```html
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Hello</title>
</head>
<body>
    <p>Hello {{.}}</p>
</body>
</html>
```

#### 解析和渲染模板文件

然后我们创建一个`main.go`文件，在其中写下HTTP server端代码如下：

```go
package main

import "net/http"

import "html/template"

import "fmt"

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	t, _ := template.ParseFiles("./hello.tmpl")
	// 渲染模板
	t.Execute(w,"沙河小王子")
}
func main() {
	http.HandleFunc("/hello",sayHello)
	err := http.ListenAndServe(":8888",nil)
	if err != nil {
		fmt.Printf(":8888 启动失败%s\n",err.Error())
	}
}
```

将上面的`main.go`文件编译执行，然后使用浏览器访问`http://127.0.0.1:8888`就能看到页面上显示了“Hello 沙河小王子”。 这就是一个最简单的模板渲染的示例，Go语言模板引擎详细用法请往下阅读。

## 模板语法

### {{.}}

模板语法都包含在`{{`和`}}`中间，其中`{{.}}`中的点表示当前对象。

当我们传入一个结构体对象时，我们可以根据`.`来访问结构体的对应字段。例如：

```GO
package main

import "net/http"

import "html/template"

import "log"

type Preson struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func Test(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./hello.tmpl")
	person := Preson{
		Name: "沙河小王子",
		Age: 18,
	}
	t.Execute(w,person)
}

func LogPanic(handler http.HandlerFunc)http.HandlerFunc {
	return func(w http.ResponseWriter,r *http.Request){
		defer func() {
			if x:=recover();x!=nil {
				log.Printf("caught panic:%s",r.RemoteAddr)
			}
		}()
		handler(w,r)
	}
}

func main() {

	http.HandleFunc("/",LogPanic(Test))
	log.Printf(":8889端口启动....")
	err := http.ListenAndServe(":8889",nil)
	if err !=nil {
		panic("ListenAndServe :8889 error"+err.Error())
	}
}
```

模板文件`hello.tmpl`内容如下：

```GO
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Hello</title>
</head>
<body>
    <p>Hello 我是{{.Name}}</p>
    <p>Hello 我的年龄{{.Age}}</p>
</body>
</html>
```

同理，当我们传入的变量是map时，也可以在模板文件中通过`.`根据key来取值

```go
type Preson struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func Test(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./hello.tmpl")
	person := Preson{
		Name: "沙河小王子",
		Age: 18,
	}
	myMap  := make(map[string]Preson)
	myMap["person1"] = person
	
	t.Execute(w,myMap)
}
```

```html
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Hello</title>
</head>
<body>
    <p>Hello 我是{{.person1.Name}}</p>
    <p>Hello 我的年龄{{.person1.Age}}</p>
</body>
</html>
```

### 注释

```go
{{/* a comment */}}
注释，执行时会忽略。可以多行。注释不能嵌套，并且必须紧贴分界符始止。
```

### pipeline

`pipeline`是指产生数据的操作。比如`{{.}}`、`{{.Name}}`等。Go的模板语法中支持使用管道符号`|`链接多个命令，用法和unix下的管道类似：`|`前面的命令会将运算结果(或返回值)传递给后一个命令的最后一个位置。

**注意：**并不是只有使用了`|`才是pipeline。Go的模板语法中，`pipeline的`概念是传递数据，只要能产生数据的，都是`pipeline`。

### 变量

我们还可以在模板中声明变量，用来保存传入模板的数据或其他语句生成的结果。具体语法如下：

```template
$obj := {{.}}
```

其中`$obj`是变量的名字，在后续的代码中就可以使用该变量了。

```go
 {{/* 下面定义一个变量*/}}
 {{$v := 1000}}
 <p>自定义变量:{{$v}}</p>
```

### 移除空格

有时候我们在使用模板语法的时候会不可避免的引入一下空格或者换行符，这样模板最终渲染出来的内容可能就和我们想的不一样，这个时候可以使用`{{-`语法去除模板内容左侧的所有空白符号， 使用`-}}`去除模板内容右侧的所有空白符号。

例如：

```template
{{- .Name -}}
```

**注意：**`-`要紧挨`{{`和`}}`，同时与模板值之间需要使用空格分隔。

### 条件判断

Go模板语法中的条件判断有以下几种:

```template
{{if pipeline}} T1 {{end}}

{{if pipeline}} T1 {{else}} T0 {{end}}

{{if pipeline}} T1 {{else if pipeline}} T0 {{end}}
```

```go
{{/*条件判断相关的*/}}
  {{if $v}}
  <p>$v:{{$v}}</p>
  {{else}} 
  <p>啥都没有</p>
  {{end}}
  <p>我是:{{.person.Name}}</p>
   <p>年龄:{{.person.Age}} {{ if ge .person.Age $v }}我成年了{{else}}我未成年{{end}}</p>
```

### range

Go的模板语法中使用`range`关键字进行遍历，有以下两种写法，其中`pipeline`的值必须是数组、切片、字典或者通道。

```template
{{range pipeline}} T1 {{end}}
如果pipeline的值其长度为0，不会有任何输出

{{range pipeline}} T1 {{else}} T0 {{end}}
如果pipeline的值其长度为0，则会执行T0。
```

例如:

```go
func Test(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./hello.tmpl")
	person := Preson{
		Name: "沙河 小王子",
		Age: 18,
	}
	myMap  := make(map[string]Preson)
	myMap["person1"] = person
	
	hobbyList := []string{
		"打篮球",
		"打泰拳",
		"踢足球",
		"xxoo",
	}
	t.Execute(w,map[string]interface{}{
		"person":person,
		"myMap":myMap,
		"hobby":hobbyList,
	})
}

func LogPanic(handler http.HandlerFunc)http.HandlerFunc {
	return func(w http.ResponseWriter,r *http.Request){
		defer func() {
			if x:=recover();x!=nil {
				log.Printf("caught panic:%v",x)
			}
		}()
		handler(w,r)
	}
}
```

模板文件:

```go
{{/* range 相关*/}}
{{ range $idx,$hobby := .hobby }}
<li>{{ $idx }}-----{{ $hobby }}</li>
{{else}}
<li>啥爱好都木有</li>
{{end}}
```

### with

```template
{{with pipeline}} T1 {{end}}
如果pipeline为empty不产生输出，否则将dot设为pipeline的值并执行T1。不修改外面的dot。

{{with pipeline}} T1 {{else}} T0 {{end}}
如果pipeline为empty，不改变dot并执行T0，否则dot设为pipeline的值并执行T1。
```

例如：

```go
<p>with 相关 相应于在with 于end之前声明一个.myMap.person1的变量,作用域是with与end之间</p>
  <p>{{.myMap}}</p>
  <p>我是:{{.myMap.person1.Name}}</p>
  <p>年龄:{{.myMap.person1.Age}}</p>
  {{ with .myMap.person1}}
  <p>我是:{{ .Name }}</p>
  <p>年龄:{{ .Age }}</p>
  {{ end }}
```

### 预定义函数

执行模板时，函数从两个函数字典中查找：首先是模板函数字典，然后是全局函数字典。一般不在模板内定义函数，而是使用Funcs方法添加函数到模板里。

预定义的全局函数如下：

```template
and
    函数返回它的第一个empty参数或者最后一个参数；
    就是说"and x y"等价于"if x then y else x"；所有参数都会执行；
or
    返回第一个非empty参数或者最后一个参数；
    亦即"or x y"等价于"if x then x else y"；所有参数都会执行；
not
    返回它的单个参数的布尔值的否定
len
    返回它的参数的整数类型长度
index
    执行结果为第一个参数以剩下的参数为索引/键指向的值；
    如"index x 1 2 3"返回x[1][2][3]的值；每个被索引的主体必须是数组、切片或者字典。
print
    即fmt.Sprint
printf
    即fmt.Sprintf
println
    即fmt.Sprintln
html
    返回与其参数的文本表示形式等效的转义HTML。
    这个函数在html/template中不可用。
urlquery
    以适合嵌入到网址查询中的形式返回其参数的文本表示的转义值。
    这个函数在html/template中不可用。
js
    返回与其参数的文本表示形式等效的转义JavaScript。
call
    执行结果是调用第一个参数的返回值，该参数必须是函数类型，其余参数作为调用该函数的参数；
    如"call .X.Y 1 2"等价于go语言里的dot.X.Y(1, 2)；
    其中Y是函数类型的字段或者字典的值，或者其他类似情况；
    call的第一个参数的执行结果必须是函数类型的值（和预定义函数如print明显不同）；
    该函数类型值必须有1到2个返回值，如果有2个则后一个必须是error接口类型；
    如果有2个返回值的方法返回的error非nil，模板执行会中断并返回给调用模板执行者该错误；
```

例如:

