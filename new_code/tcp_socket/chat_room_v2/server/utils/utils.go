package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/linzhenlong/my-go-code/new_code/tcp_socket/chat_room_v2/common/message"
	"net"
)

// 将这些方法关联到结构体中
type Transfer struct {
	// 链接
	Conn net.Conn

	// 网络传输数据buffer
	Buf [8096]byte

}

// 读包
func (transfer *Transfer)ReadPkg() (msg message.Message, err error) {

	fmt.Println("读取客户端发送消息")

	_, err = transfer.Conn.Read(transfer.Buf[:4])
	if err != nil {
		return
	}

	// 根据buf[:4]转成一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(transfer.Buf[:4])

	// 根据pkgLen读取消息内容
	n, err := transfer.Conn.Read(transfer.Buf[:pkgLen])
	if uint32(n) != pkgLen || err != nil {
		return
	}
	err = json.Unmarshal(transfer.Buf[:pkgLen], &msg)
	if err != nil {
		return
	}
	return
}

// 发送数据包函数.
func (transfer *Transfer) WritePkg(data []byte) (err error) {
	// 先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))

	// 将一个int 转成字节切片
	//var buf [4]byte
	binary.BigEndian.PutUint32(transfer.Buf[:4], pkgLen)
	n, err := transfer.Conn.Write(transfer.Buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("transfer.Conn.Write(transfer.Buf[0:4]) error", err)
		return
	}

	// 发送数据本身
	writeDataLen, err := transfer.Conn.Write(data)
	if uint32(writeDataLen) != pkgLen || err != nil {
		fmt.Println("transfer.Conn.Write(data) error", err)
		return
	}
	return
}


