package dbops

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)
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
	defer stmt.Close()
	return nil
}