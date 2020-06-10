package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	gormdb, err := gorm.Open("mysql", "root:@(127.0.0.1:3306)/ms?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = gormdb

}

func test() {
	rows, _ := db.Raw("show tables").Rows()
	fmt.Println(rows)
	var tables []string
	for rows.Next() {
		var name string
		rows.Scan(&name)
		tables = append(tables, name)
	}
	fmt.Println(tables)

	rows2, _ := db.Raw("select id,product_name, product_num from product where id in(?,?)", 2, 3).Rows()
	for rows2.Next() {
		var (
			id           int64
			product_name string
			product_num  int64
		)
		rows2.Scan(&id, &product_name, &product_num)
		fmt.Println(id, product_name, product_num)
	}
}

type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	defer db.Close()
	//test()

	// 创建表，自动迁移(把结构体和数据表进行对应)
	db.AutoMigrate(&UserInfo{})

	// 创建数据行
	u1 := UserInfo{
		//ID:     1,
		Name:   "小王子",
		Gender: "男",
		Hobby:  "篮球",
	}
	db.Create(&u1)

	// 查询
	var u UserInfo
	db.First(&u, 3)
	fmt.Println(u)

	// 更新
	db.Model(&u).Update("Hobby", "橄榄球")
	db.First(&u)
	fmt.Printf("u:%#v", u)

	// 删除
	db.Delete(&u)
}
