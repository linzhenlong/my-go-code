package process

import (
	"encoding/json"
	"fmt"
	"github.com/linzhenlong/my-go-code/new_code/tcp_socket/chat_room_v2/common/message"
	"github.com/linzhenlong/my-go-code/new_code/tcp_socket/chat_room_v2/server/model"
	"github.com/linzhenlong/my-go-code/new_code/tcp_socket/chat_room_v2/server/utils"
	"net"
)

type UserProcess struct {
	Conn net.Conn

	// 增加有关字段表示，该Coon 是哪个用户的
	UserId int
}

// 编写一个函数serverProcessLogin函数,专门处理登录请求
func (userProcess *UserProcess) ServerProcessLogin(msg *message.Message) (err error) {

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

	// 到redis 中校验数据
	user, err := model.MyUserDao.Login(loginMsg.UserId, loginMsg.UserPwd)
	fmt.Println(user)
	if err != nil {

		if err == model.ERROR_USER_NOT_EXISTS {
			loginResMsg.ErrorCode = 300
		} else if err == model.ERROR_USER_PWD {
			loginResMsg.ErrorCode = 400
		} else {
			loginResMsg.ErrorCode = 500
		}
		loginResMsg.ErrorMsg = err.Error()
	} else {
		loginResMsg.ErrorCode = 200
		loginResMsg.ErrorMsg = "success"
	}

	/*// 如果用户id 为100,密码=123456 ,认为合法否则不合法
	if loginMsg.UserId == 100 && loginMsg.UserPwd == "123456" {
		// 登录成功返回200状态码
		loginResMsg.ErrorCode = 200
		loginResMsg.ErrorMsg = "success"
	} else {
		// 登录错误返回500错误码
		loginResMsg.ErrorCode = 500
		loginResMsg.ErrorMsg = "用户不存在"
	}
	*/
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

	// 创建一个transfer 实例去实现写包操作
	transfer := &utils.Transfer{
		Conn: userProcess.Conn,
	}
	err = transfer.WritePkg(responseMsgJson)
	if err != nil {
		return
	}
	return
}

func (UserProcess *UserProcess)ServerProcessRegister(msg *message.Message)(err error) {

	// 解析消息
	registerMsg := message.RegisterMsg{}
	fmt.Println(msg.Data)
	err = json.Unmarshal([]byte(msg.Data),&registerMsg)

	if err != nil {
		fmt.Println("json.Unmarshal([]byte(msg.Data),&registerRequestMsg) error = ", err)
		return
	}
	user := model.User{
		UserPwd:registerMsg.User.UserPwd,
		UserId:registerMsg.User.UserId,
		UserName:registerMsg.User.UserName,
	}

	// 完成注册
	// 实例化

	status,err := model.MyUserDao.Register(user)

	//声明一个注册信息的response
	registerResponseMsg := message.RegisterResMsg{}
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResponseMsg.ErrorCode = 500
		} else if err == model.ERROR_WRITE_REDIS {
			registerResponseMsg.ErrorCode = 501
		} else if err == model.ERROR_NO_USER_ID_PWD {
			registerResponseMsg.ErrorCode = 502
		} else {
			registerResponseMsg.ErrorCode = 400
		}
		registerResponseMsg.ErrorMsg = err.Error()
	} else {
		if status {
			registerResponseMsg.ErrorCode = 200
			registerResponseMsg.ErrorMsg = "success"
		} else {
			registerResponseMsg.ErrorCode = 300
			registerResponseMsg.ErrorMsg = "注册失败未知原因"
		}
	}
	responseDataByte , err := json.Marshal(registerResponseMsg)
	if err != nil {
		fmt.Println("json.Marshal(registerResponseMsg) error=", err)
		return
	}

	responseMsg := message.Message{}
	responseMsg.Type = message.RegisterResMsgType
	responseMsg.Data = string(responseDataByte)

	responseByte , err := json.Marshal(responseMsg)
	if err != nil {
		fmt.Println("json.Marshal(responseMsg) error=", err)
		return
	}
	// 写包
	transfer := utils.Transfer{
		Conn:UserProcess.Conn,
	}
	return transfer.WritePkg(responseByte)


}