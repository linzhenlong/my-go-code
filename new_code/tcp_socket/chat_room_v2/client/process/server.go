package process

import (
	"fmt"
	"github.com/linzhenlong/my-go-code/new_code/tcp_socket/chat_room_v2/client/utils"
	"net"
	"os"
	"time"
)
const timeLayout = "2006/01/02 15:04:05"
// 与服务端保持连接
func KeyServerConnect(conn net.Conn)  {

	transfer := &utils.Transfer{
		Conn:conn,
	}
	for {
		fmt.Fprintf(os.Stderr,"[%s]客户端:%s不停的读取服务端消息...\n",time.Now().Format(timeLayout),conn.RemoteAddr().String())
		msg ,err := transfer.ReadPkg()
		fmt.Println("KeyServerConnect error=",err)
		if err !=nil {
			fmt.Println("服务器端出错了")
			break
		}
		fmt.Println(msg)
	}
}