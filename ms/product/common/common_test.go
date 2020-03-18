package common

import (
	"testing"
)

func TestGetResultRow(t *testing.T) {
	mysql, _ := NewMysqlConn()
	row, _ := mysql.Query("select * from product where id=1")
	res := GetResultRow(row)
	t.Log(res)
}

func TestGetResultRows(t *testing.T) {
	mysql, _ := NewMysqlConn()
	rows, _ := mysql.Query("select * from product")
	res := GetResultRows(rows)
	t.Log(res)
}
