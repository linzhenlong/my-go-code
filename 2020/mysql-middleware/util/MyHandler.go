package util

import (
	"fmt"
	"log"

	"github.com/siddontang/go-mysql/client"
	"github.com/siddontang/go-mysql/mysql"
	"github.com/siddontang/go-mysql/server"
)

// MyHandler ...
type MyHandler struct {
	server.EmptyHandler // 嵌套
	conn                *client.Conn
}

// NewMyHandler ...
func NewMyHandler() MyHandler {
	conn, err := client.Connect("127.0.0.1:3306", "root", "", "ms")
	if err != nil {
		log.Fatalf("client.Connect err:%s", err.Error())
	}
	return MyHandler{conn: conn}
}

// HandleQuery ...
func (m MyHandler) HandleQuery(query string) (*mysql.Result, error) {
	/* query = replaceProductTable(query)
	log.Printf("query :%s", query) */
	res, err := m.conn.Execute("select * from product_1")
	if err != nil {
		return nil, fmt.Errorf("query:`%s` err:%s", query, err.Error())
	}
	res2, err := m.conn.Execute("select * from product_2")
	if err != nil {
		return nil, fmt.Errorf("query:`%s` err:%s", query, err.Error())
	}
	// 结果合并
	res.RowDatas = append(res.RowDatas, res2.RowDatas...)
	// 受影响行相加
	res.AffectedRows += res2.AffectedRows
	return res, nil
}

//next to see 05Sql解析学习(select)使用第三方库、获取From表名【瑞客论坛 www.ruike1.com】
