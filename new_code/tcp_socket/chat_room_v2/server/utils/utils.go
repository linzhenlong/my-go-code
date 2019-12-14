package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/linzhenlong/my-go-code/new_code/tcp_socket/chat_room_v2/common/message"
	"net"
)

// 读包
func ReadPkg(conn net.Conn) (msg message.Message, err error) {
	buf := make([]byte, 4096)
	fmt.Println("读取客户端发送消息")

	_, err = conn.Read(buf[:4])
	if err != nil {
		return
	}

	// 根据buf[:4]转成一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])

	// 根据pkgLen读取消息内容
	n, err := conn.Read(buf[:pkgLen])
	if uint32(n) != pkgLen || err != nil {
		return
	}
	err = json.Unmarshal(buf[:pkgLen], &msg)
	if err != nil {
		return
	}
	return
}

// 发送数据包函数.
func WritePkg(conn net.Conn, data []byte) (err error) {
	// 先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))

	// 将一个int 转成字节切片
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], pkgLen)
	n, err := conn.Write(buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(buf[0:4]) error", err)
		return
	}

	// 发送数据本身
	writeDataLen, err := conn.Write(data)
	if uint32(writeDataLen) != pkgLen || err != nil {
		fmt.Println("conn.Write(data) error", err)
		return
	}
	return
}


