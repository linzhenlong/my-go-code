package dbops

import (
	"database/sql"
	"log"
)

var (
	dbConn *sql.DB
	err error
)

func init() {
	dbConn, err = sql.Open("mysql", "root@tcp(127.0.0.1:3306)/liu_mei_ti?charset=utf8")
	if err !=nil {
		log.Fatalf("MYSQL Connect error:=%s",err.Error())
	}
	
}
