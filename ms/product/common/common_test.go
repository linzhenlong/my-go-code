package common

import (
	"database/sql"
	"testing"
)

type product struct {
	ID           int64  `json:"id" ms:"id" sql:"id" ` // ms为自定义标签
	ProductName  string `json:"product_name" ms:"product_name" sql:"product_name"`
	ProductNum   int64  `json:"product_num" ms:"product_num" sql:"product_num"`
	ProductImage string `json:"product_image" ms:"product_image" sql:"product_image"`
	ProductURL   string `json:"product_url" ms:"product_url" sql:"product_url"`
	CreateDate   string `json:"create_date" ms:"create_date" sql:"create_date"`
}

var (
	mysql *sql.DB
)

func TestMain(m *testing.M) {
	mysql, _ = NewMysqlConn()
	m.Run()
}
func TestGetResultRow(t *testing.T) {
	t.SkipNow()
	mysql, _ := NewMysqlConn()
	row, _ := mysql.Query("select id,product_name,product_num,product_image,product_url  from product where id=1")

	res := GetResultRow(row)
	t.Log(res)
}

func TestGetResultRows(t *testing.T) {
	t.SkipNow()
	mysql, _ := NewMysqlConn()
	rows, _ := mysql.Query("select * from product")
	res := GetResultRows(rows)
	t.Log(res)
}

func TestDataToStruct(t *testing.T) {
	t.SkipNow()
	row, _ := mysql.Query("select * from product where id=1")
	res := GetResultRow(row)
	productRes := &product{}
	DataToStructByTagSQL(res, productRes)
	t.Log(productRes)
}
