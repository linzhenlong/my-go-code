package myrpc

import (
	"encoding/gob"
	_ "fmt"
	"net"
	"testing"
	"time"
)

// 测试:服务端有一个查询用户的方法
// 客户端去调用
type User struct {
	Name string
	Age  int
}

// 服务端用于用户查询的方法
func queryUser(id int) (User, error) {
	// 造假数据
	userMap := make(map[int]User)
	userMap[0] = User{"张三", 20}
	userMap[1] = User{"ls", 21}
	userMap[2] = User{"wu", 24}
	userMap[3] = User{"zl", 19}

	user, ok := userMap[id]
	if ok {
		return user, nil
	}
	return User{}, nil
}
func TestMyPRC(t *testing.T) {
	// 给gob 注册类型
	gob.Register(User{})

	//gob.Register()
	// 1.创建服务端对象
	addr := ":6060"
	server := NewServer(addr)

	// 2.注册服务端方法
	server.Register("queryUser", queryUser)

	// 3.服务端启动监听

	go func() {
		server.Run()
	}()
	time.Sleep(time.Second * 1)
	// 客户端
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		t.Fatalf("err:%v", err)
	}
	// 客户端对象
	cli := NewClient(conn)

	var queryFn func(int) (User, error)
	cli.callRPC("queryUser", &queryFn)

	//调用函数
	user, _ := queryFn(100)
	t.Log(user)
}
