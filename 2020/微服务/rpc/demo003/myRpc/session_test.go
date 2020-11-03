package myrpc

import "testing"

import "sync"

import "net"

func TestReadWrite(t *testing.T) {
	t.SkipNow()
	addr := ":6060"
	myData := "hello Golang"
	wg := sync.WaitGroup{}
	wg.Add(2)
	// 写的协程
	go func() {
		defer wg.Done()
		listener, err := net.Listen("tcp", addr)
		if err != nil {
			t.Log(err)
			return
		}
		conn, err := listener.Accept()
		if err != nil {
			t.Log(err)
			return
		}
		session := Session{
			conn: conn,
		}
		session.Write([]byte(myData))
	}()
	// 读的协程
	go func() {
		defer wg.Done()
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			t.Log(err)
			return
		}
		session := Session{
			conn: conn,
		}
		data, err := session.Read()
		if err != nil {
			t.Log(err)
			return
		}
		t.Log(string(data))
	}()
	wg.Wait()
}
