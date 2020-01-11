package dbops

import (
	"log"
	"database/sql"
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