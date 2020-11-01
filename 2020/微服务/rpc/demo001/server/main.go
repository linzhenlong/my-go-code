package main

import "net/rpc"

import "log"

import "net/http"

// rpc 计算矩形的面积和周长

// Rect ...
// 1.定义一个结构体，用于绑定方法
type Rect struct {
}

// Params ...
// 2.声明参数的结构体
type Params struct {
	Width  int // 长
	Hegiht int // 宽
}

// Area 面积
func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Width * p.Hegiht
	return nil
}

// ZhouChang 周长
func (r *Rect) ZhouChang(p Params, ret *int) error {
	*ret = (p.Width + p.Hegiht) * 2
	return nil
}

func main() {
	// 1. 注册服务
	err := rpc.Register(&Rect{})
	if err != nil {
		log.Println(err)
		return
	}
	// 2.把服务绑定到http协议上
	rpc.HandleHTTP()

	// 3.服务端的监听
	http.ListenAndServe(":6060", nil)
}
