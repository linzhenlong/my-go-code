package datamodels

// User 模型.
type User struct {
	ID       int64  `json:"id" gorm:"id"`
	NickName string `json:"nick_name" gorm:"nick_name"`
	UserName string `json:"user_name" gorm:"user_name"`
	Password string `json:"-" gorm:"password"`
}

// TableName 表名.
func (u *User) TableName() string {
	return "ms_user"
}
