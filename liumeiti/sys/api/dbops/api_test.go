package dbops

import "testing"



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
	dbConn.Exec("truncate user")
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
		t.Errorf("Error of DeleteUser error=%v",err)
	}
}

func TestReGetUser(t *testing.T)  {
	pwd, err := GetUserCredential("linzl")
	if err != nil {
		t.Errorf("Error of GetUserCredential error=%v",err)
	}
	t.Log(pwd)
	if len(pwd) != 0 {
		t.Errorf("Error of TestDeleteUser error=%v",pwd)
	}
}