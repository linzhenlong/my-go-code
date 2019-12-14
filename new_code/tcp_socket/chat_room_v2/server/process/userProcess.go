package process

import (
	"encoding/json"
	"github.com/linzhenlong/my-go-code/new_code/tcp_socket/chat_room_v2/common/message"
	"github.com/linzhenlong/my-go-code/new_code/tcp_socket/chat_room_v2/server/utils"
	"net"
)

// 编写一个函数serverProcessLogin函数,专门处理登录请求
func ServerProcessLogin(conn net.Conn, msg *message.Message) (err error) {

	// 核心代码
	// 1. 先从msg中取出msg.Data,并直接反序列化成LoginMsg
	loginMsg := message.LoginMsg{}

	err = json.Unmarshal([]byte(msg.Data), &loginMsg)
	if err != nil {
		return
	}

	// 声明一个responseMsg
	var responseMsg message.Message
	responseMsg.Type = message.LoginResMsgType

	// 再声明一个 LongResMsg
	loginResMsg := message.LoginResMsg{}

	// 如果用户id 为100,密码=123456 ,认为合法否则不合法
	if loginMsg.UserId == 100 && loginMsg.UserPwd == "123456" {
		// 登录成功返回200状态码
		loginResMsg.ErrorCode = 200
		loginResMsg.ErrorMsg = "success"
	} else {
		// 登录错误返回500错误码
		loginResMsg.ErrorCode = 500
		loginResMsg.ErrorMsg = "用户不存在"
	}

	// 序列化一下
	loginResMsgJson, err := json.Marshal(loginResMsg)

	if err != nil {
		return
	}

	// 赋值给responseMsg
	responseMsg.Data = string(loginResMsgJson)

	// 对responseMsg 序列化
	responseMsgJson, err := json.Marshal(responseMsg)
	if err != nil {
		return
	}

	// 发送responseMsgJson数据，将他封装成到一个writePkg()函数
	err = utils.WritePkg(conn, responseMsgJson)
	return
}
