package myrpc

import (
	"log"
	"net"
	"reflect"
)

// Server 服务端结构体
type Server struct {
	// 服务端监听地址
	addr string
	// 维护服务端函数名到函数值的映射的map
	funcs map[string]reflect.Value
}

// NewServer ...
func NewServer(addr string) *Server {
	return &Server{
		addr:  addr,
		funcs: make(map[string]reflect.Value),
	}
}

// Register 服务注册方法
func (s *Server) Register(funcName string, f interface{}) {
	// 判断一下map是否已经影射了这个方法
	if _, ok := s.funcs[funcName]; ok {
		return
	}
	// 注册方法
	s.funcs[funcName] = reflect.ValueOf(f)
}

// Run 服务端实现的方法...
func (s *Server) Run() {
	// 1.监听
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Println(err)
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			return
		}
		// 2.创建一个会话,服务端session
		servSession := NewSession(conn)

		// 3.使用rpc方式读取客户端数据
		b, err := servSession.Read()
		if err != nil {
			log.Println(err)
			return
		}
		// 解码客户端数据
		rpcData, err := decode(b)
		if err != nil {
			log.Println(err)
			return
		}
		// 5.根据读取的数据的name,调用具体方法
		f, ok := s.funcs[rpcData.Name]
		if !ok {
			log.Println("函数不存在")
			return
		}
		// 6.调用函数,解析Args
		// 利用反射的Call(args)方法调函数传参,args 必须是[]ref.Value类型的
		myArgs := make([]reflect.Value, 0, len(rpcData.Args))
		for _, arg := range rpcData.Args {
			myArgs = append(myArgs, reflect.ValueOf(arg))
		}
		// 反射调用函数，返回一个切片类型的Value
		resp := f.Call(myArgs)

		// 7.返回结果
		// 需要遍历，可以在编码，在返回
		respArgs := make([]interface{}, 0, len(resp))
		for _, respVal := range resp {
			respArgs = append(respArgs, respVal.Interface())
		}
		// 需要构建rpcData
		respRPCData := RPCData{
			Name: rpcData.Name,
			Args: respArgs,
		}
		bytes, err := encode(respRPCData)
		if err != nil {
			log.Println(err)
			return
		}
		//log.Println("resp",string(bytes))

		// 8.写出数据
		err = servSession.Write(bytes)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
