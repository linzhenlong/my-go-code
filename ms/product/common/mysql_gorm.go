package common

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// NewGorm 初始化.
func NewGorm() (db *gorm.DB, err error) {
	db, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/ms?charset=utf8")
	return
}
