package myrpc

import (
	"log"
	"net"
	"reflect"
)

// Client 客户端..
type Client struct {
	conn net.Conn
}

// NewClient ...
func NewClient(conn net.Conn) *Client {
	return &Client{
		conn: conn,
	}
}

func (c *Client) callRPC(rpcName string, fPtr interface{}) {
	// 通过反射获取函数原型
	fn := reflect.ValueOf(fPtr).Elem()

	// 客户端逻辑的实现
	f := func(args []reflect.Value) []reflect.Value {
		// 创建客户端会话
		cliSession := NewSession(c.conn)

		// 构建请求参数
		inArgs := make([]interface{}, 0, len(args))
		for _, arg := range args {
			inArgs = append(inArgs, arg.Interface())
		}
		// 构建rpc数据
		requesRPCData := RPCData{
			Name: rpcName,
			Args: inArgs,
		}
		// 编码
		requestBytes, err := encode(requesRPCData)
		if err != nil {
			log.Println(err)
			return nil
		}
		// 写数据到服务端(请求服务端)
		err = cliSession.Write(requestBytes)
		if err != nil {
			log.Println(err)
			return nil
		}
		// 接收服务端返回的结果
		respBytes, err := cliSession.Read()
		if err != nil {
			log.Println(err)
			return nil
		}
		// 解码服务端数据
		respData, err := decode(respBytes)
		if err != nil {
			log.Println(err)
			return nil
		}
		// 处理服务端返回的结果
		outArgs := make([]reflect.Value, 0, len(respData.Args))
		for i, arg := range respData.Args {
			// 特殊处理结果为nil的情况
			if arg == nil {
				//reflect.Zero():返回某了类型的零值Value
				// Out():函数输出的参数类型
				// 得到具体第几个位置的参数的零值
				outArgs = append(outArgs, reflect.Zero(fn.Type().Out(i)))
			} else {
				outArgs = append(outArgs, reflect.ValueOf(arg))
			}
		}
		return outArgs
	}
	// 函数到原型调用的关键
	// 参数1:函数原型，是Type 类型
	// 参数2:回调函数
	value := reflect.MakeFunc(fn.Type(), f)
	fn.Set(value)
}
