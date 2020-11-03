package myrpc

import (
	"encoding/binary"
	"io"
	"net"
)

// Session ...定义一个会话的结构体.
type Session struct {
	conn net.Conn
}

// NewSession 构造方法.
func NewSession(conn net.Conn) *Session {
	return &Session{
		conn: conn,
	}
}

// Write 向连接中写数据
func (s *Session) Write(data []byte) error {
	// 1.记录数据长度
	// 2.记录数据
	// 4字节的头长度 head uint32+data
	buf := make([]byte, 4+len(data))

	// 给buf 前4个位置写数据的长度,作为请求头
	binary.BigEndian.PutUint32(buf[:4], uint32(len(data)))

	// 写入真正的数据，
	copy(buf[4:], data)
	_, err := s.conn.Write(buf)
	if err != nil {
		return err
	}
	/* if n != len(data) {
		return errors.New("丢包了")
	} */
	return nil
}

// 连接中读数据
func (s *Session) Read() ([]byte, error) {

	// 1.读取头部数据
	header := make([]byte, 4)
	// 按照指定长度读取数据
	_, err := io.ReadFull(s.conn, header)
	if err != nil {
		return nil, err
	}
	// 2.读取数据，通过header 解出数据长度
	dataLen := binary.BigEndian.Uint32(header)
	data := make([]byte, dataLen)
	_, err = io.ReadFull(s.conn, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
