package myrpc

import "encoding/gob"

import "bytes"

// RPCData 定义rpc交互数据格式
type RPCData struct {
	Name string
	Args []interface{}
}

// 编码
func encode(data RPCData) ([]byte, error) {
	// gob 进行编码
	var buf bytes.Buffer
	// 编码器
	encoder := gob.NewEncoder(&buf)
	if err := encoder.Encode(data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// 解码
func decode(b []byte)  (*RPCData,error) {
	buf := bytes.NewBuffer(b)
	// 解码器
	decoder := gob.NewDecoder(buf)
	data := &RPCData{}
	// 对数据解码
	decoder.Decode(data)
	return data,nil
}
