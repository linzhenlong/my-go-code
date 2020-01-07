package dbops

import "testing"
import "strconv"
import "time"


var tempUUID string

// 初始化测试用例
// 1.清空数据
// 2.跑测试用例
// 3.在清空数据
func TestMain(m *testing.M)  {
	clearTables()
	m.Run()
	clearTables()
}


// 清空数据
func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

// 1.添加用户
// 2.获取用户
// 3.删除用户
// 4.在重新获取用户
func TestUserWorkFlow(t *testing.T)  {
	t.Run("AddUser",TestAddUser)
	t.Run("GetUser",TestGetUser)
	t.Run("DeleteUser",TestDeleteUser)
	t.Run("ReGetUser",TestReGetUser)
}

func TestAddUser(t *testing.T)  {
	err := AddUserCredential("linzl", "123456")
	if err != nil {
		t.Errorf("Error of TestAddUser error:%v",err)
	}
}

func TestGetUser(t *testing.T) {
	pwd, err := GetUserCredential("linzl")
	if err != nil {
		t.Errorf("Error of GetUserCredential error=%v",err)
	}
	if pwd != "123456" {
		t.Fatalf("GetUserCredential pwd error,期望值%s,实际值:%s", "123456",pwd)
	}
}

func TestDeleteUser(t *testing.T) {
	//t.SkipNow()
	err := DeleteUser("linzl", "123456")
	if err != nil {
		t.Fatalf("Error of DeleteUser error=%v",err)
	}
}

func TestReGetUser(t *testing.T)  {
	pwd, err := GetUserCredential("linzl")
	if err != nil {
		t.Fatalf("Error of GetUserCredential error=%v",err)
	}
	t.Log(pwd)
	if len(pwd) != 0 {
		t.Fatalf("Error of TestDeleteUser error=%v",pwd)
	}
}

func TestVidoFlow(t *testing.T)  {
	t.Run("addVideo", TestAddVideo)
	t.Run("getVideo", TestGetVideoIno)
	t.Run("delVideo", TestDeleteVideo)
	t.Run("reGetVideo",TestGetVideoIno)
}

func TestAddVideo(t *testing.T)  {
	res, err := AddNewVideo(1,"MY-Video")
	if err != nil {
		t.Fatalf("AddNewVideo error=%s", err.Error())
	}
	if res == nil {
		t.Fatalf("AddNewVideo res error%v", res)
	}
	tempUUID = res.UUID
	t.Log("TestAddVideo",res,tempUUID)
}

func TestGetVideoIno(t *testing.T)  {
	t.Log("TestGetVideoIno uuid==>",tempUUID)
	res, err := GetVidoInfo(tempUUID)
	if err != nil {
		t.Fatalf("TestGetVideoIno error=%s", err.Error())
	}
	t.Logf("TestGetVideoIno:%v",res)
}

func TestDeleteVideo(t *testing.T)  {
	err := DeleteVideo(tempUUID)
	if err != nil {
		t.Logf("TestDeleteVideo error=%s", err.Error())
	}
	t.Log("TestDeleteVideo succ")
}

func TestComments(t *testing.T) {
	clearTables()
	// 添加一个用户,联查用
	t.Run("TestCommentsAddUser", TestAddUser)
	// 添加一个视频
	t.Run("TestCommentsAddVideo", TestAddVideo)
	// 添加一条评论
	t.Run("TestAddComment", TestAddComment)
	t.Run("TestAddComment", TestAddComment)
	t.Run("TestAddComment", TestAddComment)
	t.Run("TestAddComment", TestAddComment)
	t.Run("TestAddComment", TestAddComment)
	t.Run("TestAddComment", TestAddComment)
	t.Run("TestAddComment", TestAddComment)
	t.Run("TestAddComment", TestAddComment)
	t.Run("TestAddComment", TestAddComment)
	t.Run("TestAddComment", TestAddComment)
	t.Run("TestAddComment", TestAddComment)	
	t.Run("TestListComments", TestListComments)	
}

func TestAddComment(t *testing.T) {
	err := AddNewComments(tempUUID, 1, "TestAddComment"+strconv.Itoa(int(time.Now().UnixNano())))
	if err != nil {
		t.Fatalf("TestAddComment error =%s", err.Error())
	}
	t.Log("TestAddComment succ")
}

func TestListComments(t *testing.T) {
	res, err := CommentsList(tempUUID, 0, 20)
	if err != nil {
		t.Logf("TestListComments error=%s",err)
	}
	for i,v := range res {
		t.Log(i,v)
	}
}