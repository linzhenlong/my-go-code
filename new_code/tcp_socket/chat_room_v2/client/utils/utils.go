package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/linzhenlong/my-go-code/new_code/tcp_socket/chat_room_v2/common/message"
	"net"
	"os"
	"time"
)

const timeLayout = "2006/01/02 15:04:05"

type Transfer struct {
	Conn net.Conn
	Buf [8096]byte``
}

func (transfer *Transfer)ReadPkg()(msg message.Message, err error)  {
	// 获取包长度
	fmt.Fprintln(os.Stderr,"【"+time.Now().Format(timeLayout)+"】客户端读包....")
	_, err = transfer.Conn.Read(transfer.Buf[:4])
	if err !=nil {
		return
	}

	// 将buf[:4]转成uint32类型的
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(transfer.Buf[0:4])

	// 读取真正的数据包
	n, err := transfer.Conn.Read(transfer.Buf[0:pkgLen])
	if err != nil {
		fmt.Fprintln(os.Stderr,"【"+time.Now().Format(timeLayout)+"】transfer.Conn.Read(transfer.Buf[0:pkgLen] error",err)
		return
	}
	if uint32(n) != pkgLen {
		fmt.Fprintln(os.Stderr,"【"+time.Now().Format(timeLayout)+"】transfer.Conn.Read(transfer.Buf[0:pkgLen] 包长度不一致")
		return
	}
	err = json.Unmarshal(transfer.Buf[0:pkgLen],&msg)
	if err !=nil {
		fmt.Fprintln(os.Stderr,"【"+time.Now().Format(timeLayout)+"】Unmarshal(transfer.Buf[0:pkgLen],&msg) 错误error=",err)
		return
	}
	return
}

func (transfer *Transfer)WritePkg(data []byte)(err error) {

	// 先发送一个head 长度
	// 需要将int 转换为byte切片
	fmt.Fprintln(os.Stderr,"【"+time.Now().Format(timeLayout)+"】客户端发包....")
	pkgLen := uint32(len(data))
	binary.BigEndian.PutUint32(transfer.Buf[0:4],pkgLen)

	n,err := transfer.Conn.Write(transfer.Buf[0:4])
	if n!=4 || err !=nil {
		fmt.Fprintln(os.Stderr,"【"+time.Now().Format(timeLayout)+"】transfer.Conn.Write(transfer.Buf[0:4]) error=",
			err)
		return
	}

	// 发送数据
	// 发送数据本身
	writeDataLen, err := transfer.Conn.Write(data)
	if uint32(writeDataLen) != pkgLen || err != nil {
		fmt.Fprintln(os.Stderr,"【"+time.Now().Format(timeLayout)+"】transfer.Conn.Write(data) error=",
			err)
		return
	}
	return
}