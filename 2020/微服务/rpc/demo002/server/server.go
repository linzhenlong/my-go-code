package main

import "net/http"

import "net/rpc"

import "log"

import "errors"

// Calc ...
type Calc struct {
}

// Params ... 参数
type Params struct {
	X int
	Y int
}

// Resp 返回值...
type Resp struct {
	Chengji int
	Shang   float64
	Yushu   int
}

// Cheng ...
func (c *Calc) Cheng(p Params, res *Resp) error {
	res.Chengji = p.X * p.Y
	return nil
}

// Jia ...
func (c *Calc) Jia(p Params, ret *int) error {
	*ret = p.X + p.Y
	return nil
}

// Jian ...
func (c *Calc) Jian(p Params, ret *int) error {
	*ret = p.X - p.Y
	return nil
}

// Chu ...
func (c *Calc) Chu(p Params, res *Resp) error {
	if p.Y == 0 {
		return errors.New("除数不能为零")
	}
	res.Shang = float64(p.X / p.Y)
	return nil
}

// All ...
func (c *Calc) All(p Params, res *Resp) error {
	if p.Y == 0 {
		return errors.New("除数不能为零")
	}
	*res = Resp{
		Chengji: p.X * p.Y,
		Shang:   float64(p.X / p.Y),
		Yushu:   p.X % p.Y,
	}
	return nil
}

func main() {
	err := rpc.Register(&Calc{})
	if err != nil {
		log.Println(err)
		return
	}
	rpc.HandleHTTP()

	err = http.ListenAndServe(":6060", nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("rpc 启动成功")
}
