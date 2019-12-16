package process

import (
	"encoding/json"
	"fmt"
	"github.com/linzhenlong/my-go-code/new_code/tcp_socket/chat_room_v2/client/utils"
	"github.com/linzhenlong/my-go-code/new_code/tcp_socket/chat_room_v2/common/message"
	"net"
	"os"
)

type UserProcess struct {

}

func (userProcess *UserProcess)Login(userId int ,userPassword string) (err error) {

	// 链接到服务器端
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		return  err
	}
	// 延迟关闭连接
	defer conn.Close()

	// 通过conn发送消息
	loginMsg := message.LoginMsg{
		UserId:userId,
		UserPwd:userPassword,
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

	// 实例化
	transfer := &utils.Transfer{
		Conn:conn,
	}

	//// 先发送一个消息长度
	//// 先获取requestData长度，然后长度转成，表示长度的[]byte
	//
	//// 发送的数据包长度
	//var msgPackageLen = uint32(len(requestData))
	//
	//var msgBytes [4]byte
	//binary.BigEndian.PutUint32(msgBytes[:], msgPackageLen)
	//
	//// 实例化
	//transfer := &utils.Transfer{
	//	Conn:conn,
	//}
	//err = transfer.WritePkg(msgBytes[:])
	//if err !=nil {
	//	fmt.Println("conn.Write 失败 ","err=",err)
	//	return nil
	//}
	//
	//fmt.Println("客户端发送消息长度=",len(requestData))



	// 发送消息本身
	err = transfer.WritePkg(requestData)
	if err !=nil {
		fmt.Println("conn.Write 失败 err=",err)
		return nil
	}

	// 获取服务端响应的消息

	responseMsg , err := transfer.ReadPkg()
	if err != nil {
		fmt.Println("utils.ReadPkg(conn) error=",err)
		return err
	}

	// 将responseMsg的data部分反序列成LoginResMsg
	loginResMsg := message.LoginResMsg{}
	err = json.Unmarshal([]byte(responseMsg.Data), &loginResMsg)
	if err != nil {
		fmt.Println("responseMsg.Data 反序列化出错，error=", err)
		return err
	}

	if loginResMsg.ErrorCode == 200 {

		// 起一个协程保持与服务端通讯
		go KeyServerConnect(conn)
		for {
			LoginSuccessView()
		}



	} else {
		fmt.Println(loginResMsg.ErrorMsg)
	}
	return
}

func LoginSuccessView() {

		fmt.Println("----用户登录界面---")
		fmt.Println("----1.显示在线用户列表---")
		fmt.Println("----2.发送消息---")
		fmt.Println("----3.信息列表---")
		fmt.Println("----4.退出系统---")
		fmt.Println("请选择:(1-4):")
		var key int
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("----1.显示在线用户列表---")

		case 2:
			fmt.Println("----2.发送消息---")

		case 3:
			fmt.Println("----3.信息列表---")

		case 4:
			fmt.Println("你退出了系统")
			os.Exit(0)
		default:
			fmt.Println("输入有误，请输入(1-4)")
		}

}