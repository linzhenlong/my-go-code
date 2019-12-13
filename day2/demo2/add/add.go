package add

import (
     _ "go_dev/day2/demo2/test" // 下划线的作用,导入这个包但是没有用到这个包里的变量或是方法
)

var Name string = "张三"
var Age int  = 18

func init() {
    Name = "张三"
    Age = 18
}
