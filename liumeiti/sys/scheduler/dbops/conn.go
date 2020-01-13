package dbops

import (
	"database/sql"
	"fmt"
	"log"
)

var (
	dbConn *sql.DB
	err error
)

func init() {
	dbConn, err = sql.Open("mysql", "root@tcp(127.0.0.1:3306)/liu_mei_ti?charset=utf8")
	fmt.Println(err)
	if err !=nil {
		log.Fatalf("MYSQL Connect111 error:=%s",err.Error())
	}
	
}
