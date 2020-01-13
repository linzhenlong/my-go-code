package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // 需要引入这个包,并且_掉，不然 conn 会err:sql: unknown driver "mysql" (forgotten import?)
	"log"
)

// ReadVideoDeletionRecord 读取待删除视频id.
func ReadVideoDeletionRecord(count int)([]string, error) {
	log.Print("ReadVideoDeletionRecord start....")
	stmtOut, err := dbConn.Prepare("SELECT video_uuid FROM video_del_rec limit ?")
	var ids []string
	if err != nil {
		log.Printf("ReadVideoDeletionRecord mysql select error:=%s\n", err.Error())
		return ids, err
	}
	rows, err := stmtOut.Query(count)
	if err !=nil && err != sql.ErrNoRows {
		log.Printf("ReadVideoDeletionRecord mysql query error:%s",err.Error())
		return ids, err
	}
	if err == sql.ErrNoRows {
		return ids, nil
	}
	for rows.Next() {
		var videoUUID string
		err = rows.Scan(&videoUUID)
		if err != nil {
			continue
		}
		ids = append(ids, videoUUID)
	}
	defer stmtOut.Close()
	return ids, err
}

//DelVideoDeletionRecord 删除视频的相关记录.
func DelVideoDeletionRecord(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM video_del_rec WHERE video_uuid=?")
	if err != nil {
		log.Printf("DelVideoDeletionRecord MYSQL prepare error:%s", err.Error())
		return err
	}
	defer stmtDel.Close()

	_, err = stmtDel.Exec(vid)
	if err != nil {
		log.Printf("DelVideoDeletionRecord Delete error:%s", err.Error())
		return err
	}
	return nil
}

// AddVideoDeletionRecord 添加删除记录.
func AddVideoDeletionRecord(vid string) error {
	stmt, err := dbConn.Prepare("insert into video_del_rec(video_uuid) values(?)")
	if err != nil {
		log.Printf("AddVideoDeletionRecord Prepare error:%s",err.Error())
		return err
	}
	_,err = stmt.Exec(vid)
	if err != nil {
		log.Printf("AddVideoDeletionRecord sql Exec error:%s", err.Error())
		return err
	}
	return nil
}