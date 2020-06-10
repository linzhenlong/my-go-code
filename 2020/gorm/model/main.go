package main

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// User 用户结构体,模型.
type User struct {
	gorm.Model          // 内嵌gorm.Model 模型
	Name         string // 零值类型
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"column:user_email;type:varchar(100);unique_index"` // 指定列表名是user_email
	Role         string  `gorm:"size:255"`                                         // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"`                                  // 设置MemberNumber 唯一且不允许为空
	Num          int     `gorm:"AUTO_INCREMENT"`                                   // 自增
	Address      string  `gorm:"index:idx_addr"`
	IgnoreMe     int     `gorm:"-"` //忽略本字段
}

// Monster 结构体 model.
type Monster struct {
	MonsterID int64 `gorm:"primary_key"` // 指定主键
	Name      string
	Age       int
}

// TableName 修改默认表名.
func (Monster) TableName() string {
	return "monster"
}

func main() {
	db, err := gorm.Open("mysql", "root:@(127.0.0.1:3306)/ms?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// 禁用表名的复数形式
	db.SingularTable(true)

	// 对默认表名添加前缀.
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "ms_" + defaultTableName
	}

	// 创建表
	err = db.AutoMigrate(&User{}).Error
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Monster{})

	// db.Table() 指定表名,相当于使用user结构体创建一个表名为user_bak 的表.
	db.Table("user_bak").AutoMigrate(&User{})

}
