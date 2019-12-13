package service

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/linzhenlong/mygo/go_dev/new_code/tcp_socket/chat_room/client/model"
	"github.com/linzhenlong/mygo/go_dev/new_code/tcp_socket/chat_room/common/message"
	"net"
)

type UserService struct {
	users *[]model.User

	userId int

	userPassword string

}

func NewUserService(userId int,userPassword string) *UserService  {
	return &UserService{
		userId:userId,
		userPassword:userPassword,
	}
}

func (u *UserService) Login() error  {
	fmt.Printf("userid=%d,password=%s\n", u.userId,u.userPassword)

	// 链接到服务器端
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		return  err
	}
	// 延迟关闭连接
	defer conn.Close()

	// 通过conn发送消息
	loginMsg := message.LoginMsg{
		UserId:u.userId,
		UserPwd:u.userPassword,
	}
	requestMsgData, err := json.Marshal(loginMsg)
	if err != nil {
		return  err
	}
	requestMsg := message.Message{
		Type:message.LoginMsgType,
		Data:string(requestMsgData),
	}

	requestData , err := json.Marshal(requestMsg)
	if err != nil {
		return  err
	}

	// 先发送一个消息长度
	// 先获取requestData长度，然后长度转成，表示长度的[]byte

	// 发送的数据包长度
	var msgPackageLen = uint32(len(requestData))

	var msgBytes [4]byte
	binary.BigEndian.PutUint32(msgBytes[:], msgPackageLen)

	requestLen , err := conn.Write(msgBytes[:])
	if requestLen !=4 || err !=nil {
		fmt.Println("conn.Write 失败 requestLen=",requestLen,"err=",err)
		return nil
	}

	fmt.Println("客户端发送消息长度=",len(requestData))

	// 发送消息本身
	requestLen , err = conn.Write(requestData)
	if err !=nil {
		fmt.Println("conn.Write 失败 requestLen=",requestLen,"err=",err)
		return nil
	}

	return nil
}