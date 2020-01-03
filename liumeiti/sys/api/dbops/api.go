package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/linzhenlong/my-go-code/liumeiti/sys/api/defs"
	"github.com/linzhenlong/my-go-code/liumeiti/sys/api/utils"
	_ "log"
	"time"
)

// AddUserCredential 添加用户.
func AddUserCredential(loginName, pwd string) error {
	stmtIns, err := dbConn.Prepare("INSERT into users(`login_name`,`pwd`) VALUES (?, ?)")
	if err != nil {
		return err
	}

	//pwdMd5 := utils.GenMd5(pwd)
	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

// GetUserCredential 获取用户.
func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("SELECT pwd FROM `users` WHERE `login_name`= ?")
	if err != nil {
		//MyLog.Printf("GetUserCredential error=%s",err.Error())
		return "", err
	}
	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	defer stmtOut.Close()
	return pwd, nil

}

// DeleteUser 删除用户 .
func DeleteUser(loginName, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM `users`  WHERE `login_name`= ? AND `pwd`=? limit 1")
	if err != nil {
		//MyLog.Printf("GetUserCredential error=%s",err.Error())
		return err
	}
	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtDel.Close()
	return nil
}

// AddNewVideo 添加视频.
func AddNewVideo(aid int, name string) (videoInfo *defs.VideoInfo, err error) {
	uuid, err := utils.NewUUID()
	if err != nil {
		return
	}
	createTime := time.Now().Format("2006-01-02 15:04:05")

	sqlTemplate := "INSERT INTO `video_info`(uuid,name,display_ctime,author_id) VALUES(?, ?, ?, ?)"
	stmtIns, err := dbConn.Prepare(sqlTemplate)
	if err != nil {
		return
	}
	sqlRes, err := stmtIns.Exec(uuid, name, createTime, aid)

	if err != nil {
		return
	}
	id, err := sqlRes.LastInsertId()
	if err != nil {
		return
	}
	defer stmtIns.Close()
	videoInfo = &defs.VideoInfo{
		ID:           int(id),
		AuthorID:     aid,
		Name:         name,
		DisplayCtime: createTime,
		UUID:         uuid,
	}
	return
}

// GetVidoInfo 获取视频详情.
func GetVidoInfo(uuid string) (videoInfo *defs.VideoInfo, err error) {
	sqlTemplate := "SELECT id,author_id,name,display_ctime FROM `video_info` where uuid=? limit 1"
	stmtOut, err := dbConn.Prepare(sqlTemplate)
	if err != nil {
		return
	}
	var (
		id           int
		authorID     int
		name         string
		displayCtime string
	)
	err = stmtOut.QueryRow(uuid).Scan(&id, &authorID, &name, &displayCtime)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	videoInfo = &defs.VideoInfo{
		ID:           id,
		AuthorID:     authorID,
		Name:         name,
		DisplayCtime: displayCtime,
		UUID:         uuid,
	}
	defer stmtOut.Close()
	return
}

// DeleteVideo .
func DeleteVideo(uuid string) (err error) {
	sqlTemplate := "DELETE FROM `video_info` where uuid=? limit 1"
	smts, err := dbConn.Prepare(sqlTemplate)
	if err != nil {
		goto END
	}
	_, err = smts.Exec(uuid)
	if err != nil {
		goto END
	}
	defer smts.Close()
END:
	return
}

// AddNewComments 添加评论.
func AddNewComments(videoUUID string, authorID int, content string) (err error) {
	sqlTemplate := "insert into `comments` (video_uuid,author_id,content) VALUES(?, ?, ?)"

	stmtIns, err := dbConn.Prepare(sqlTemplate)
	if err != nil {
		goto END
	}
	_, err = stmtIns.Exec(videoUUID, authorID, content)
	if err != nil {
		goto END
	}
	defer stmtIns.Close()
END:
	return
}

// CommentsList 评论列表.
func CommentsList(videoUUID string, from, size int) (list []*defs.Comment, err error) {
	sqlTemplate := `select comments.*, users.login_name from comments 
					inner join users ON comments.author_id = users.id  
					where comments.video_uuid=? order by time desc limit ?, ?`

	stmtOut, err := dbConn.Prepare(sqlTemplate)
	if err != nil {
		return
	}
	rows, err := stmtOut.Query(videoUUID, from, size)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	if err == sql.ErrNoRows {
		return
	}
	for rows.Next() {
		var (
			id       int
			uuid     string
			authorID int
			content  string
			time     string
			userName string
		)
		err = rows.Scan(&id, &uuid, &authorID, &content, &time, &userName)
		if err != nil && err != sql.ErrNoRows {
			continue
		}
		comment := &defs.Comment{
			ID:        id,
			VideoUUID: uuid,
			AuthorID:  authorID,
			Content:   content,
			Time:      time,
			UserName:  userName,
		}
		list = append(list, comment)
	}
	err = rows.Err()
	if err != nil && err != sql.ErrNoRows {
		return
	}
	defer stmtOut.Close()

	return
}
