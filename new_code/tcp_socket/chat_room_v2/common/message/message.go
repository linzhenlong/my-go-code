package message

const (
	LoginMsgType = "loginRequest"
	LoginResMsgType = "loginResponse"
	RegisterMsgType = "register"
	RegisterResMsgType = "RegisterResMsg"
)

type Message struct {
	Type string  `json:"type"`// 消息类型
	Data string  `json:"data"`// 消息体
}


// 登录的消息
type LoginMsg struct {
	UserId int `json:"user_id"`
	UserPwd string `json:"user_pwd"`
	//UserName string `json:"user_name"`
}

// 登录服务端返回消息
type LoginResMsg struct {
	ErrorCode int  `json:"error_code"`// 返回状态码，200成功,等等...
	ErrorMsg string `json:"error_msg"`
}

// 注册信息.
type RegisterMsg struct {
	User User `json:"user"`
}

// 注册的响应信息
type RegisterResMsg struct {
	ErrorCode int  `json:"error_code"`// 返回状态码，200成功,等等...
	ErrorMsg string `json:"error_msg"`
}
