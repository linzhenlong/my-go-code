package common

import (
	"database/sql"
	"log"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
)

// NewMysqlConn 创建数据库连接.
func NewMysqlConn() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/ms?charset=utf8")
	return
}

// GetResultRow 获取结果集中的一条数据
func GetResultRow(rows *sql.Rows) map[string]string {
	resultRow := make(map[string]string)
	colums, _ := rows.Columns()
	log.Printf("colums:%v", colums)
	scanArgs := make([]interface{}, len(colums))
	values := make([]interface{}, len(colums))
	// 为scan 时做准备
	for j := range values {
		scanArgs[j] = &values[j]
	}
	log.Printf("scanArgs:%v", scanArgs)
	log.Printf("values:%v", values)
	for rows.Next() {
		rows.Scan(scanArgs...)
		for i, v := range values {
			if v != nil {
				log.Printf("v:%v", v)
				log.Printf("k:%v", colums[i])
				log.Printf("vtype:%v", reflect.TypeOf(v))
				resultRow[colums[i]] = string(v.([]byte))
			}
		}
	}
	return resultRow
}

// GetResultRows 获取所有含数据.
func GetResultRows(rows *sql.Rows) map[int]map[string]string {
	resultRows := make(map[int]map[string]string)
	colums, _ := rows.Columns()
	log.Printf("colums:%v", colums)
	values := make([][]byte, len(colums))
	scanArgs := make([]interface{}, len(colums))

	for j := range values {
		scanArgs[j] = &values[j]
	}
	i := 0
	for rows.Next() {
		rows.Scan(scanArgs...)
		log.Printf("scanArgs-%d:%v", i, scanArgs)
		log.Printf("values-%d:%v", i, values)
		row := make(map[string]string)
		for k, v := range values {
			key := colums[k]
			row[key] = string(v)
			log.Printf("row[%s]=%v", key, string(v))
		}

		resultRows[i] = row
		i++
	}
	return resultRows
}
