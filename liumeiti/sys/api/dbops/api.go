package dbops

import (
	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/linzhenlong/my-go-code/liumeiti/sys/api/utils"
	_ "log"
)


func AddUserCredential(loginName ,pwd string) error  {
	stmtIns, err := dbConn.Prepare("INSERT into user(`login_name`,`pwd`) VALUES (?, ?)")
	if err != nil {
		return err
	}

	//pwdMd5 := utils.GenMd5(pwd)
	stmtIns.Exec(loginName, pwd)
	stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string)(string, error)  {
	stmtOut ,err := dbConn.Prepare("SELECT pwd FROM `user` WHERE `login_name`= ?")
	if err != nil {
		//MyLog.Printf("GetUserCredential error=%s",err.Error())
		return "", err
	}
	var pwd string
	stmtOut.QueryRow(loginName).Scan(&pwd)
	stmtOut.Close()
	return pwd,nil

}

func DeleteUser(loginName ,pwd string) error {
	stmtDel , err := dbConn.Prepare("DELETE FROM `user`  WHERE `login_name`= ? AND `pwd`=? limit 1")
	if err != nil {
		//MyLog.Printf("GetUserCredential error=%s",err.Error())
		return err
	}
	stmtDel.Exec(loginName, pwd)
	stmtDel.Close()
	return nil
}

// 该看视频3-11