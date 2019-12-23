package process

import "fmt"

// 因为userMgr实例在服务器端有且只有一个
// 因为在很多地方都会使用他，因为把他做成全局变量

var userMgr *UserMgr

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}


//初始化userMgr
func init()  {
	userMgr = &UserMgr{
		onlineUsers:make(map[int]*UserProcess,1024),
	}
}

// 添加在线用户
func (UserMgr *UserMgr)AddOnlineUser(userProcess *UserProcess) {
	UserMgr.onlineUsers[userProcess.UserId] = userProcess
}

func (UserMgr *UserMgr)DeleteOnlineUser(userProcess *UserProcess)  {

	// 判断该用户是否在线，如果在线将他从map删除
	if _,ok := UserMgr.onlineUsers[userProcess.UserId];ok {
		delete(UserMgr.onlineUsers, userProcess.UserId)
	}

}

// 获取当前所有在线用户
func (UserMgr *UserMgr)GetAllOnlineUsers() map[int]*UserProcess  {
	return UserMgr.onlineUsers
}

func (UserMgr *UserMgr)GetOnlineUserById(userId int)(userProcess *UserProcess ,err error) {
	userProcess, ok := UserMgr.onlineUsers[userId]
	if !ok {
		err = fmt.Errorf("当前用户%d,不在线",userId)
		goto END
	}
	
	END:
	return
}